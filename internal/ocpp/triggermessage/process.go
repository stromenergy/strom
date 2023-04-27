package triggermessage

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *TriggerMessage) ProcessConf(client *ws.Client, messageCall types.MessageCall) {
	_, err := unmarshalTriggerMessageConf(messageCall.Payload)

	if err != nil {
		util.LogError("STR035: Error unmarshaling TriggerMessageConf", err)
		callError := types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeFORMATIONVIOLATION, "", types.NoError{})
		callError.Send(client)
		return
	}
}
