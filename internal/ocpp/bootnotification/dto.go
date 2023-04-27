package bootnotification

import (
	"encoding/json"

	"github.com/stromenergy/strom/internal/ocpp/types"
)

type BootNotificationReq struct {
	ChargePointModel        string  `json:"chargePointModel"`
	ChargePointVendor       string  `json:"chargePointVendor"`
	ChargeBoxSerialNumber   *string `json:"chargeBoxSerialNumber"`
	ChargePointSerialNumber *string `json:"chargePointSerialNumber"`
	FirmwareVersion         *string `json:"firmwareVersion"`
	Iccid                   *string `json:"iccid"`
	Imsi                    *string `json:"imsi"`
	MeterSerialNumber       *string `json:"meterSerialNumber"`
	MeterType               *string `json:"meterType"`
}

type BootNotificationConf struct {
	CurrentTime types.OcppTime           `json:"currentTime"`
	Interval    int                      `json:"interval"`
	Status      types.RegistrationStatus `json:"status"`
}

func unmarshalBootNotificationReq(payload []byte) (*BootNotificationReq, error) {
	bootNotificationReq := &BootNotificationReq{}

	if err := json.Unmarshal(payload, bootNotificationReq); err != nil {
		return nil, err
	}

	return bootNotificationReq, nil
}
