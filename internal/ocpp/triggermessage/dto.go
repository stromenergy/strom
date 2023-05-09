package triggermessage

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/ocpp/types"
)

type TriggerMessageReq struct {
	MessageTrigger types.MessageTrigger `json:"messageTrigger"`
	ConnectorID    *int32               `json:"connectorId,omitempty"`
}

type TriggerMessageConf struct {
	Status types.TriggerMessageStatus `json:"status"`
}

func UnmarshalTriggerMessageConf(payload interface{}) (*TriggerMessageConf, error) {
	triggerMessageConf := &TriggerMessageConf{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, triggerMessageConf); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return triggerMessageConf, nil
}
