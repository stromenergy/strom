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

func UnmarshalChangeAvailabilityConf(payload interface{}) (*ChangeAvailabilityConf, error) {
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

type ChangeConfigurationReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ChangeConfigurationConf struct {
	Status types.ConfigurationStatus `json:"status"`
}

func UnmarshalChangeConfigurationConf(payload interface{}) (*ChangeConfigurationConf, error) {
	changeConfigurationConf := &ChangeConfigurationConf{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, changeConfigurationConf); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return changeConfigurationConf, nil
}

type ClearCacheReq struct{}

type ClearCacheConf struct {
	Status types.ClearCacheStatus `json:"status"`
}

func UnmarshalClearCacheReq(payload interface{}) (*ClearCacheReq, error) {
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

type KeyValue struct {
	Key      string  `json:"key"`
	Readonly bool    `json:"readonly"`
	Value    *string `json:"value"`
}

type GetConfigurationReq struct {
	Key *[]string `json:"key,omitempty"`
}

type GetConfigurationConf struct {
	ConfigurationKey *[]KeyValue `json:"configurationKey,omitempty"`
	UnknownKey       *[]string   `json:"unknownKey,omitempty"`
}

func UnmarshalGetConfigurationConf(payload interface{}) (*GetConfigurationConf, error) {
	getConfigurationConf := &GetConfigurationConf{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, getConfigurationConf); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return getConfigurationConf, nil
}

type ResetReq struct {
	Type types.ResetType `json:"type"`
}

type ResetConf struct {
	Status types.ResetStatus `json:"status"`
}

func UnmarshalResetReq(payload interface{}) (*ResetReq, error) {
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

func UnmarshalUnlockConnectorConf(payload interface{}) (*UnlockConnectorConf, error) {
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
