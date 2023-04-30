package transaction

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/metervalue"
	"github.com/stromenergy/strom/internal/ocpp/types"
)

type IDTagInfo struct {
	ExpiryDate  *types.OcppTime           `json:"expiryDate,omitempty"`
	ParentIDTag *string                   `json:"parentIdTag,omitempty"`
	Status      types.AuthorizationStatus `json:"status"`
}

type StartTransactionReq struct {
	ConnectorID   int32          `json:"connectorId"`
	IDTag         string         `json:"idTag"`
	MeterStart    int32          `json:"meterStart"`
	ReservationID *int64         `json:"reservationId,omitempty"`
	Timestamp     types.OcppTime `json:"timestamp"`
}

type StartTransactionConf struct {
	IDTagInfo     IDTagInfo `json:"idTagInfo"`
	TransactionID int64     `json:"transactionId"`
}

func unmarshalStartTransactionReq(payload interface{}) (*StartTransactionReq, error) {
	startTransactionReq := &StartTransactionReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, startTransactionReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return startTransactionReq, nil
}

type StopTransactionReq struct {
	IDTag           *string                   `json:"idTag,omitempty"`
	MeterStop       int32                     `json:"meterStop"`
	Timestamp       types.OcppTime            `json:"timestamp"`
	TransactionID   int64                     `json:"transactionId"`
	Reason          *db.TransactionStopReason `json:"reason,omitempty"`
	TransactionData []metervalue.MeteredValue `json:"transactionData,omitempty"`
}

type StopTransactionConf struct {
	IDTagInfo *IDTagInfo `json:"idTagInfo,omitempty"`
}

func unmarshalStopTransactionReq(payload interface{}) (*StopTransactionReq, error) {
	stopTransactionReq := &StopTransactionReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, stopTransactionReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return stopTransactionReq, nil
}
