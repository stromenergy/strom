package management

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Management) SendUnlockConnectorReq(client *ws.Client, connectorId int32) (string, <-chan types.Message, error) {
	cancelReservationReq := UnlockConnectorReq{
		ConnectorID: connectorId,
	}

	return s.call.Send(client, types.CallActionUnlockConnector, cancelReservationReq)
}
