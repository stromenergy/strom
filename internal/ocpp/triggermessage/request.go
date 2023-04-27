package triggermessage

import (
	"github.com/google/uuid"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func Request(client *ws.Client, messageTrigger types.MessageTrigger, connectorId *int) {
	triggerMessageReq := TriggerMessageReq{
		MessageTrigger: messageTrigger,
		ConnectorID: connectorId,
	}

	messageCall := types.NewMessageCall(uuid.NewString(), types.ActionTRIGGERMESSAGE, triggerMessageReq)
	messageCall.Send(client)
}
