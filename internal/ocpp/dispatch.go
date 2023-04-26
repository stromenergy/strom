package ocpp

import (
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Ocpp) Broadcast(message []byte) {
	s.broadcast <- message
}

func (s *Ocpp) Register(client *ws.Client) {
	s.register <- client
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
		case message := <-s.broadcast:
			for client := range s.clients {
				select {
				case client.Send <- message:
				default:
					s.remove(client)
				}
			}
		}
	}
}

func (s *Ocpp) remove(client *ws.Client) {
	delete(s.clients, client)
	close(client.Send)
}