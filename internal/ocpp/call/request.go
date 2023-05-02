package call

import (
	"github.com/google/uuid"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Call) Send(client *ws.Client, action types.CallAction, payload interface{}) (string, <-chan types.Message) {
	uniqueID := s.getUniqueID()
	s.channels[uniqueID] = make(chan types.Message)

	message := types.NewMessageCall(uniqueID, action, payload)
	message.Send(client)

	return uniqueID, s.channels[uniqueID]
}

func (s *Call) getUniqueID() string {
	for {
		uniqueID := uuid.NewString()

		if _, ok := s.channels[uniqueID]; !ok {
			return uniqueID
		}
	}
}
