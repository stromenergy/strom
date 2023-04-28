package statusnotification

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
)

type StatusNotificationReq struct {
	ConnectorID     int32                   `json:"connectorId"`
	ErrorCode       db.ChargePointErrorCode `json:"errorCode"`
	Status          db.ChargePointStatus    `json:"status"`
	Info            *string                 `json:"info,omitempty"`
	Timestamp       *types.OcppTime         `json:"timestamp,omitempty"`
	VendorID        *string                 `json:"vendorId,omitempty"`
	VendorErrorCode *string                 `json:"vendorErrorCode,omitempty"`
}

type StatusNotificationConf struct {
}

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
