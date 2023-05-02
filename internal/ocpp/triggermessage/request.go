package triggermessage

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *TriggerMessage) Request(client *ws.Client, chargePointId int64, messageTrigger types.MessageTrigger, connectorId *int32) (string, <-chan types.Message) {
	triggerMessageReq := TriggerMessageReq{
		MessageTrigger: messageTrigger,
		ConnectorID:    connectorId,
	}

	return s.call.Send(client, types.CallActionTriggerMessage, triggerMessageReq)
}
