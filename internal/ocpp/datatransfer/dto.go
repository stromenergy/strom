package datatransfer

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/ocpp/types"
)

type DataTransferReq struct {
	VendorID  string  `json:"vendorId"`
	MessageID *string `json:"messageId,omitempty"`
	Data      *string `json:"data,omitempty"`
}

type DataTransferConf struct {
	Status types.DataTransferStatus `json:"status"`
	Data   *string                  `json:"data,omitempty"`
}

func unmarshalDataTransferReq(payload interface{}) (*DataTransferReq, error) {
	dataTransferReq := &DataTransferReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, dataTransferReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return dataTransferReq, nil
}

func unmarshalDataTransferConf(payload interface{}) (*DataTransferConf, error) {
	dataTransferConf := &DataTransferConf{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, dataTransferConf); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return dataTransferConf, nil
}
