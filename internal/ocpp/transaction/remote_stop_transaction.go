package transaction

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Transaction) SendRemoteStopTransactionReq(client *ws.Client, transactionID int64) (string, <-chan types.Message, error) {
	remoteStopTransactionReq := RemoteStopTransactionReq{
		TransactionID: transactionID,
	}

	return s.call.Send(client, types.CallActionRemoteStartTransaction, remoteStopTransactionReq)
}
