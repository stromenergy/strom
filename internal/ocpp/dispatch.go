package ocpp

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/internal/ws"
)

const (
	OCPP_VERSION = "ocpp1.6"
)

func (s *Ocpp) CheckOrigin(r *http.Request) bool {
	protocol := r.Header.Get("Sec-Websocket-Protocol")

	return strings.Contains(protocol, OCPP_VERSION)
}

func (s *Ocpp) Packet(packet *ws.Packet) {
	s.inbox <- packet
}

func (s *Ocpp) Register(client *ws.Client, params gin.Params) {
	if id, ok := params.Get("id"); ok {
		client.ID = id
		log.Debug().Msgf("Registering charge point: %s", id)
	}

	s.register <- client
}

func (s *Ocpp) Subprotocols() []string {
	return []string{OCPP_VERSION}
}

func (s *Ocpp) Unregister(client *ws.Client) {
	s.unregister <- client
}

func (s *Ocpp) dispatch() {
	for {
		select {
		case client := <-s.register:
			s.clients[client] = true
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				s.remove(client)
			}
		case packet := <-s.inbox:
			s.processPacket(packet)
		}
	}
}

func (s *Ocpp) remove(client *ws.Client) {
	delete(s.clients, client)
	client.CloseQueue()
}