package notification

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/db"
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

type StatusNotificationReq struct {
	ConnectorID     int32                   `json:"connectorId"`
	ErrorCode       db.ChargePointErrorCode `json:"errorCode"`
	Status          db.ChargePointStatus    `json:"status"`
	Info            *string                 `json:"info,omitempty"`
	Timestamp       *types.OcppTime         `json:"timestamp,omitempty"`
	VendorID        *string                 `json:"vendorId,omitempty"`
	VendorErrorCode *string                 `json:"vendorErrorCode,omitempty"`
}

type StatusNotificationConf struct{}

func unmarshalStatusNotificationReq(payload interface{}) (*StatusNotificationReq, error) {
	statusNotificationReq := &StatusNotificationReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, statusNotificationReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return statusNotificationReq, nil
}
