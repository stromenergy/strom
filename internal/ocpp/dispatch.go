package ocpp

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/util"
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
	ctx := context.Background()

	if chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID); err == nil && chargePoint.Status == db.ChargePointStatusOnline {
		updateChargePointParams := param.NewUpdateChargePointParams(chargePoint)
		updateChargePointParams.Status = db.ChargePointStatusOffline

		if _, err := s.repository.UpdateChargePoint(ctx, updateChargePointParams); err != nil {
			util.LogError("STR086: Error updating charge point", err)
		}
	}

	delete(s.clients, client)
	client.CloseQueue()
}
