package bootnotification

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/ocpp/types"
)

type BootNotificationReq struct {
	ChargePointModel        string  `json:"chargePointModel"`
	ChargePointVendor       string  `json:"chargePointVendor"`
	ChargeBoxSerialNumber   *string `json:"chargeBoxSerialNumber,omitempty"`
	ChargePointSerialNumber *string `json:"chargePointSerialNumber,omitempty"`
	FirmwareVersion         *string `json:"firmwareVersion,omitempty"`
	Iccid                   *string `json:"iccid,omitempty"`
	Imsi                    *string `json:"imsi,omitempty"`
	MeterSerialNumber       *string `json:"meterSerialNumber,omitempty"`
	MeterType               *string `json:"meterType,omitempty"`
}

type BootNotificationConf struct {
	CurrentTime types.OcppTime           `json:"currentTime"`
	Interval    int                      `json:"interval"`
	Status      types.RegistrationStatus `json:"status"`
}

func unmarshalBootNotificationReq(payload interface{}) (*BootNotificationReq, error) {
	bootNotificationReq := &BootNotificationReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, bootNotificationReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return bootNotificationReq, nil
}
