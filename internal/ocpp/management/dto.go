package management

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/ocpp/types"
)

type ChangeAvailabilityReq struct {
	ConnectorID int32                  `json:"connectorId"`
	Type        types.AvailabilityType `json:"type"`
}

type ChangeAvailabilityConf struct {
	Status types.AvailabilityStatus `json:"status"`
}

func unmarshalChangeAvailabilityConf(payload interface{}) (*ChangeAvailabilityConf, error) {
	chargeAvailabilityConf := &ChangeAvailabilityConf{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, chargeAvailabilityConf); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return chargeAvailabilityConf, nil
}

type ClearCacheReq struct{}

type ClearCacheConf struct {
	Status types.ClearCacheStatus `json:"status"`
}

func unmarshalClearCacheReq(payload interface{}) (*ClearCacheReq, error) {
	clearCacheReq := &ClearCacheReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, clearCacheReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return clearCacheReq, nil
}

type ResetReq struct {
	Type types.ResetType `json:"type"`
}

type ResetConf struct {
	Status types.ResetStatus `json:"status"`
}

func unmarshalResetReq(payload interface{}) (*ResetReq, error) {
	resetReq := &ResetReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, resetReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return resetReq, nil
}

type UnlockConnectorReq struct {
	ConnectorID int32 `json:"connectorId"`
}

type UnlockConnectorConf struct {
	Status types.UnlockStatus `json:"status"`
}

func unmarshalUnlockConnectorConf(payload interface{}) (*UnlockConnectorConf, error) {
	unlockConnectorConf := &UnlockConnectorConf{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, unlockConnectorConf); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return unlockConnectorConf, nil
}
