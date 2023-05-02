package call

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Call) Receive(client *ws.Client, message types.Message) {
	if channel, ok := s.channels[message.UniqueID]; ok {
		channel <- message

		s.Remove(message.UniqueID)
	}
}
