package ocpp

import (
	"encoding/json"

	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Ocpp) processPacket(packet *ws.Packet) {
	messageCall := types.MessageCall{}

	if err := json.Unmarshal(packet.Message, &messageCall); err != nil {
		util.LogError("STR031: Error unmarshaling message", err)
		return
	}

	switch messageCall.Action {
	case types.ActionBOOTNOTIFICATION:
		s.bootNotification.ProcessReq(packet.Client, messageCall)
	case types.ActionHEARTBEAT:
		s.heartbeat.ProcessReq(packet.Client, messageCall)
	case types.ActionTRIGGERMESSAGE:
		s.triggerMessage.ProcessConf(packet.Client, messageCall)
	default:
		callError := types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeNOTSUPPORTED, "", types.NoError{})
		callError.Send(packet.Client)
	}
}
