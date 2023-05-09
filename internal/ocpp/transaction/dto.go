package transaction

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/authorization"
	"github.com/stromenergy/strom/internal/ocpp/metervalue"
	"github.com/stromenergy/strom/internal/ocpp/types"
)

type ChargingProfile struct{}

type RemoteStartTransactionReq struct {
	ConnectorID     *int32           `json:"connectorId,omitempty"`
	IDTag           string           `json:"idTag"`
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty"`
}

type RemoteStartTransactionConf struct {
	Status types.RemoteStartStopStatus `json:"status"`
}

func UnmarshalRemoteStartTranformationConf(payload interface{}) (*RemoteStartTransactionReq, error) {
	remoteStartTransactionReq := &RemoteStartTransactionReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, remoteStartTransactionReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return remoteStartTransactionReq, nil
}

type RemoteStopTransactionReq struct {
	TransactionID int64 `json:"transactionId"`
}

type RemoteStopTransactionConf struct {
	Status types.RemoteStartStopStatus `json:"status"`
}

func UnmarshalRemoteStopTranformationConf(payload interface{}) (*RemoteStopTransactionReq, error) {
	remoteStopTransactionReq := &RemoteStopTransactionReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, remoteStopTransactionReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return remoteStopTransactionReq, nil
}

type StartTransactionReq struct {
	ConnectorID   int32          `json:"connectorId"`
	IDTag         string         `json:"idTag"`
	MeterStart    int32          `json:"meterStart"`
	ReservationID *int64         `json:"reservationId,omitempty"`
	Timestamp     types.OcppTime `json:"timestamp"`
}

type StartTransactionConf struct {
	IDTagInfo     authorization.IDTagInfo `json:"idTagInfo"`
	TransactionID int64                   `json:"transactionId"`
}

func UnmarshalStartTransactionReq(payload interface{}) (*StartTransactionReq, error) {
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
	IDTagInfo *authorization.IDTagInfo `json:"idTagInfo,omitempty"`
}

func UnmarshalStopTransactionReq(payload interface{}) (*StopTransactionReq, error) {
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
