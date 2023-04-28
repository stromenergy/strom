package triggermessage

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *TriggerMessage) Request(client *ws.Client, chargePointId int64, messageTrigger types.MessageTrigger, connectorId *int32) {
	if call, err := s.call.Create(chargePointId, db.CallActionTriggerMessage); err == nil {
		triggerMessageReq := TriggerMessageReq{
			MessageTrigger: messageTrigger,
			ConnectorID:    connectorId,
		}

		message := types.NewMessageCall(call.ReqID, db.CallActionTriggerMessage, triggerMessageReq)
		message.Send(client)
	}
}
