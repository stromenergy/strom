package transaction

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Transaction) SendRemoteStartTransactionReq(client *ws.Client, connectorId *int32, idTag string, chargingProfile *ChargingProfile) (string, <-chan types.Message) {
	remoteStartTransactionReq := RemoteStartTransactionReq{
		ConnectorID: connectorId,
		IDTag: idTag,
		ChargingProfile: chargingProfile,
	}

	return s.call.Send(client, types.CallActionRemoteStartTransaction, remoteStartTransactionReq)
}
