package ocpp

import (
	"encoding/json"

	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Ocpp) processPacket(packet *ws.Packet) {
	var callResult *types.MessageCallResult
	var callError *types.MessageCallError

	clientID := packet.Client.ID
	messageCall := types.MessageCall{}

	if err := json.Unmarshal(packet.Message, &messageCall); err != nil {
		util.LogError("STR028: Error unmarshaling message", err)
		return
	}

	switch messageCall.Action {
	case types.ActionBOOTNOTIFICATION:
		callResult, callError = s.bootNotification.ProcessReq(clientID, messageCall)
	case types.ActionHEARTBEAT:
		callResult, callError = s.heartbeat.ProcessReq(clientID, messageCall)
	default:
		callError = types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeNOTSUPPORTED, "", types.NoError{})
	}

	s.sendResponse(packet.Client, callResult, callError)
}

func (s *Ocpp) sendResponse(client *ws.Client, callResult *types.MessageCallResult, callError *types.MessageCallError) {
	if callResult != nil {
		bytes, err := json.Marshal(callResult)

		if err != nil {
			util.LogError("STR029: Error marshaling result response", err)
			return
		}

		client.Send(bytes)
	} else if callError != nil {
		bytes, err := json.Marshal(callError)

		if err != nil {
			util.LogError("STR030: Error marshaling error response", err)
			return
		}

		client.Send(bytes)
	}
}
