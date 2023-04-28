package bootnotification

import (
	"context"

	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/triggermessage"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)


func (s *BootNotification) ProcessReq(client *ws.Client, messageCall types.MessageCall) {
	bootNotificationReq, err := unmarshalBootNotificationReq(messageCall.Payload)

	if err != nil {
		util.LogError("STR032: Error unmarshaling BootNotificationReq", err)
		callError := types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeFORMATIONVIOLATION, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err == nil {
		// Update existing charge point
		updateChargePointParams := param.NewUpdateChargePointParams(chargePoint)
		updateChargePointParams.Model = bootNotificationReq.ChargePointModel
		updateChargePointParams.Vendor = bootNotificationReq.ChargePointVendor
		updateChargePointParams.SerialNumber = util.SqlNullString(bootNotificationReq.ChargePointSerialNumber)
		updateChargePointParams.FirmwareVerion = util.SqlNullString(bootNotificationReq.FirmwareVersion)
		updateChargePointParams.ModemIccid = util.SqlNullString(bootNotificationReq.Iccid)
		updateChargePointParams.ModemImsi = util.SqlNullString(bootNotificationReq.Imsi)
		updateChargePointParams.MeterSerialNumber = util.SqlNullString(bootNotificationReq.MeterSerialNumber)
		updateChargePointParams.MeterType = util.SqlNullString(bootNotificationReq.MeterType)

		_, err = s.repository.UpdateChargePoint(ctx, updateChargePointParams)

		if err != nil {
			util.LogError("STR033: Error updating charge point", err)
			callError := types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeINTERNALERROR, "", types.NoError{})
			callError.Send(client)
			return
		}
	} else {
		// Create new charge point
		createChargePointParams := createChargePointParams(client.ID, bootNotificationReq)

		_, err = s.repository.CreateChargePoint(ctx, createChargePointParams)

		if err != nil {
			util.LogError("STR034: Error creating charge point", err)
			callError := types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeINTERNALERROR, "", types.NoError{})
			callError.Send(client)
			return
		}

		// TODO: Create a nostr account
	}

	bootNotificationConf := BootNotificationConf{
		CurrentTime: types.NewOcppTime(nil),
		Interval:    900,
		Status:      types.RegistrationStatusACCEPTED,
	}

	callResult := types.NewMessageCallResult(messageCall.UniqueID, bootNotificationConf)
	callResult.Send(client)

	triggermessage.Request(client, types.MessageTriggerSTATUSNOTIFICATION, nil)
	triggermessage.Request(client, types.MessageTriggerMETERVALUES, nil)
	
	// TODO: Notify UI of changes
}
