package datatransfer

import (
	"context"

	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *DataTransfer) DataTransferReq(client *ws.Client, message types.Message) {
	dataTransferReq, err := UnmarshalDataTransferReq(message.Payload)

	if err != nil {
		util.LogError("STR056: Error unmarshaling DataTransferReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR057: Charge point not found", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	key := getHandlerKey(dataTransferReq.VendorID, dataTransferReq.MessageID)
	dataTransferConf := DataTransferConf{
		Status: types.DataTransferStatusUnknownVendorId,
	}

	if handler, ok := s.handlers[key]; ok {
		status, data := handler.DataTransferReq(chargePoint, dataTransferReq)
		dataTransferConf.Status = status
		dataTransferConf.Data = data
	}

	callResult := types.NewMessageCallResult(message.UniqueID, dataTransferConf)
	callResult.Send(client)
}

func (s *DataTransfer) SendDataTransferReq(client *ws.Client, vendorId string, messageID, data *string) (string, <-chan types.Message, error) {
	dataTransferReq := DataTransferReq{
		VendorID:  vendorId,
		MessageID: messageID,
		Data:      data,
	}

	return s.call.Send(client, types.CallActionDataTransfer, dataTransferReq)
}
