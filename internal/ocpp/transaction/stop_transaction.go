package transaction

import (
	"context"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Transaction) StopTransactionReq(client *ws.Client, message types.Message) {
	stopTransactionReq, err := unmarshalStopTransactionReq(message.Payload)

	if err != nil {
		util.LogError("STR052: Error unmarshaling StopTransactionReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	transaction, err := s.repository.GetTransaction(ctx, stopTransactionReq.TransactionID)

	if err != nil {
		util.LogError("STR053: Transaction not found", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	updateTransactionParams := param.NewUpdateTransactionParams(transaction)
	updateTransactionParams.MeterStop = util.SqlNullInt32(stopTransactionReq.MeterStop)
	updateTransactionParams.Reason = defaultReason(stopTransactionReq.Reason)
	updateTransactionParams.Status = db.TransactionStatusStopped
	updateTransactionParams.StopTimestamp = util.SqlNullTime(stopTransactionReq.Timestamp.Time())

	_, err = s.repository.UpdateTransaction(ctx, updateTransactionParams)

	if stopTransactionReq.TransactionData != nil {
		if chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID); err == nil {
			s.meterValue.ProcessMeterValues(ctx, chargePoint.ID, transaction.ConnectorID, &transaction.ID, stopTransactionReq.TransactionData)
		}
	}

	if err != nil {
		util.LogError("STR054: Error updating transaction", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	stopTransactionConf := StopTransactionConf{}

	if stopTransactionReq.IDTag != nil {
		idTagInfo := s.authorization.GetIDTagInfo(ctx, *stopTransactionReq.IDTag, db.AuthorizationStatusExpired)
		stopTransactionConf.IDTagInfo = &idTagInfo
	}

	callResult := types.NewMessageCallResult(message.UniqueID, stopTransactionConf)
	callResult.Send(client)

	// TODO: Notify UI of changes
}
