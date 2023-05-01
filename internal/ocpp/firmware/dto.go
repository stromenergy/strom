package firmware

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/ocpp/types"
)

type FirmwareStatusNotificationReq struct {
	Status types.FirmwareStatus `json:"status"`
}

type FirmwareStatusNotificationConf struct{}

func unmarshalFirmwareStatusNotificationReq(payload interface{}) (*FirmwareStatusNotificationReq, error) {
	firmwareStatusNotificationReq := &FirmwareStatusNotificationReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, firmwareStatusNotificationReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return firmwareStatusNotificationReq, nil
}
