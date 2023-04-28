package triggermessage

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *TriggerMessage) ProcessConf(client *ws.Client, message types.Message) {
	_, err := unmarshalTriggerMessageConf(message.Payload)

	if err != nil {
		util.LogError("STR035: Error unmarshaling TriggerMessageConf", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}
}
