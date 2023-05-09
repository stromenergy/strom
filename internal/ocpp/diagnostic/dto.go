package diagnostic

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/ocpp/types"
)

type DiagnosticsStatusNotificationReq struct {
	Status types.DiagnosticsStatus `json:"status"`
}

type DiagnosticsStatusNotificationConf struct{}

func UnmarshalDiagnosticsStatusNotificationReq(payload interface{}) (*DiagnosticsStatusNotificationReq, error) {
	diagnosticsStatusNotificationReq := &DiagnosticsStatusNotificationReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, diagnosticsStatusNotificationReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return diagnosticsStatusNotificationReq, nil
}
