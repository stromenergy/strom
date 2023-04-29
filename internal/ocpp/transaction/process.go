package transaction

import (
	"context"

	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Transaction) StartTransactionReq(client *ws.Client, message types.Message) {
	startTransactionReq, err := unmarshalStartTransactionReq(message.Payload)

	if err != nil {
		util.LogError("STR046: Error unmarshaling StartTransactionReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR047: Charge point not found", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	createTransactionParams := createTransactionParams(chargePoint.ID, startTransactionReq)
	transaction, err := s.repository.CreateTransaction(ctx, createTransactionParams)

	if err != nil {
		util.LogError("STR048: Error creating transaction", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	// TODO: Authenticate IDTag against tags table
	starTransactionConf := StartTransactionConf{
		IDTagInfo: IDTagInfo{
			Status: types.AuthorizationStatusAccepted,
		},
		TransactionID: transaction.ID,
	}

	callResult := types.NewMessageCallResult(message.UniqueID, starTransactionConf)
	callResult.Send(client)

	// TODO: Notify UI of changes
}
