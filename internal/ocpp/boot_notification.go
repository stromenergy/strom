package ocpp

import (
	"context"
	"encoding/json"
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/util"
)

type BootNotificationReq struct {
	ChargeBoxSerialNumber   *string `json:"chargeBoxSerialNumber"`
	ChargePointModel        string  `json:"chargePointModel"`
	ChargePointSerialNumber *string `json:"chargePointSerialNumber"`
	ChargePointVendor       string  `json:"chargePointVendor"`
	FirmwareVersion         *string `json:"firmwareVersion"`
	Iccid                   *string `json:"iccid"`
	Imsi                    *string `json:"imsi"`
	MeterSerialNumber       *string `json:"meterSerialNumber"`
	MeterType               *string `json:"meterType"`
}

type BootNotificationConf struct {
	CurrentTime OcppTime           `json:"currentTime"`
	Interval    int                `json:"interval"`
	Status      RegistrationStatus `json:"status"`
}

func (s *Ocpp) bootNotification(clientID string, messageCall MessageCall) (*MessageCallResult, *MessageCallError) {
	// TODO unmarshal data and update db
	bootNotificationReq, err := unmarshalBootNotificationReq(messageCall.Payload)

	if err != nil {
		util.LogError("STR031: Error unmarshaling BootNotificationReq", err)
		return nil, NewMessageCallError(messageCall.UniqueID, ErrorCodeFORMATIONVIOLATION, "", NoError{})
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
			return nil, NewMessageCallError(messageCall.UniqueID, ErrorCodeINTERNALERROR, "", NoError{})
		}
	} else {
		// Create new charge point
		createChargePointParams := createChargePointParams(clientID, bootNotificationReq)

		_, err = s.repository.CreateChargePoint(ctx, createChargePointParams)

		if err != nil {
			util.LogError("STR033: Error creating charge point", err)
			return nil, NewMessageCallError(messageCall.UniqueID, ErrorCodeINTERNALERROR, "", NoError{})
		}
	}

	bootNotificationConf := BootNotificationConf{
		CurrentTime: NewOcppTime(nil),
		Interval:    900,
		Status:      RegistrationStatusACCEPTED,
	}

	return NewMessageCallResult(messageCall.UniqueID, bootNotificationConf), nil
}

func createChargePointParams(identity string, bootNotificationReq *BootNotificationReq) db.CreateChargePointParams {
	return db.CreateChargePointParams{
		Identity:          identity,
		Model:             bootNotificationReq.ChargePointModel,
		Vendor:            bootNotificationReq.ChargePointVendor,
		SerialNumber:      util.SqlNullString(bootNotificationReq.ChargePointSerialNumber),
		FirmwareVerion:    util.SqlNullString(bootNotificationReq.FirmwareVersion),
		ModemIccid:        util.SqlNullString(bootNotificationReq.Iccid),
		ModemImsi:         util.SqlNullString(bootNotificationReq.Imsi),
		MeterSerialNumber: util.SqlNullString(bootNotificationReq.MeterSerialNumber),
		MeterType:         util.SqlNullString(bootNotificationReq.MeterType),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func unmarshalBootNotificationReq(payload []byte) (*BootNotificationReq, error) {
	bootNotificationReq := &BootNotificationReq{}

	if err := json.Unmarshal(payload, bootNotificationReq); err != nil {
		return nil, err
	}

	return bootNotificationReq, nil
}
