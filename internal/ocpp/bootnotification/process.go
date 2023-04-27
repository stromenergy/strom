package bootnotification

import (
	"context"

	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
)


func (s *BootNotification) ProcessReq(clientID string, messageCall types.MessageCall) (*types.MessageCallResult, *types.MessageCallError) {
	// TODO unmarshal data and update db
	bootNotificationReq, err := unmarshalBootNotificationReq(messageCall.Payload)

	if err != nil {
		util.LogError("STR031: Error unmarshaling BootNotificationReq", err)
		return nil, types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeFORMATIONVIOLATION, "", types.NoError{})
	}

	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, clientID)

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
			util.LogError("STR032: Error updating charge point", err)
			return nil, types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeINTERNALERROR, "", types.NoError{})
		}
	} else {
		// Create new charge point
		createChargePointParams := createChargePointParams(clientID, bootNotificationReq)

		_, err = s.repository.CreateChargePoint(ctx, createChargePointParams)

		if err != nil {
			util.LogError("STR033: Error creating charge point", err)
			return nil, types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeINTERNALERROR, "", types.NoError{})
		}

		// TODO: Create a nostr account
	}

	// TODO: Notify UI of changes

	bootNotificationConf := BootNotificationConf{
		CurrentTime: types.NewOcppTime(nil),
		Interval:    900,
		Status:      types.RegistrationStatusACCEPTED,
	}

	return types.NewMessageCallResult(messageCall.UniqueID, bootNotificationConf), nil
}
