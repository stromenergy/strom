package metervalue

import (
	"context"

	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *MeterValue) MeterValueReq(client *ws.Client, message types.Message) {
	meterValueReq, err := UnmarshalMeterValueReq(message.Payload)

	if err != nil {
		util.LogError("STR049: Error unmarshaling MeterValueReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR050: Charge point not found", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	s.ProcessMeterValues(ctx, chargePoint.ID, meterValueReq.ConnectorID, meterValueReq.TransactionID, meterValueReq.MeterValue)

	callResult := types.NewMessageCallResult(message.UniqueID, MeterValueConf{})
	callResult.Send(client)

	// TODO: Notify UI of changes
}

func (s *MeterValue) ProcessMeterValues(ctx context.Context, chargePointID int64, connectorID int32, transactionID *int64, meterValues []MeteredValue) {
	for _, meteredValue := range meterValues {
		for _, sampledValue := range meteredValue.SampledValue {
			createMeterValueParams := createMeterValueParams(chargePointID, connectorID, transactionID, meteredValue, sampledValue)
			
			if _, err := s.repository.CreateMeterValue(ctx, createMeterValueParams); err != nil {
				util.LogError("STR051: Error creating meter value", err)
			}
		}
	}
}
