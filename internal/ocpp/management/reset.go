package management

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Management) SendResetReq(client *ws.Client, resetType types.ResetType) (string, <-chan types.Message, error) {
	resetReq := ResetReq{
		Type: resetType,
	}

	return s.call.Send(client, types.CallActionReset, resetReq)
}
