package metervalue

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
)

type SampledValue struct {
	Value     string                  `json:"value"`
	Context   *db.MeterReadingContext `json:"context,omitempty"`
	Format    *db.MeterValueFormat    `json:"format,omitempty"`
	Measurand *db.MeterMeasurand      `json:"measurand,omitempty"`
	Phase     *db.MeterPhase          `json:"phase,omitempty"`
	Location  *db.MeterLocation       `json:"location,omitempty"`
	Unit      *db.MeterUnitOfMeasure  `json:"unit,omitempty"`
}

type MeteredValue struct {
	Timestamp    types.OcppTime `json:"timestamp"`
	SampledValue []SampledValue `json:"sampledValue"`
}

type MeterValueReq struct {
	ConnectorID   int32          `json:"connectorId"`
	TransactionID *int64         `json:"transactionId,omitempty"`
	MeterValue    []MeteredValue `json:"meterValue"`
}

type MeterValueConf struct{}

func UnmarshalMeterValueReq(payload interface{}) (*MeterValueReq, error) {
	meterValueReq := &MeterValueReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, meterValueReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return meterValueReq, nil
}
