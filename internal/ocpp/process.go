package ocpp

import (
	"encoding/json"

	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Ocpp) processPacket(packet *ws.Packet) {
	var callResult *MessageCallResult
	var callError *MessageCallError

	clientID := packet.Client.ID
	messageCall := MessageCall{}

	if err := json.Unmarshal(packet.Message, &messageCall); err != nil {
		util.LogError("STR028: Error unmarshaling message", err)
		return
	}

	switch messageCall.Action {
	case ActionBOOTNOTIFICATION:
		callResult, callError = s.bootNotification(clientID, messageCall)
	default:
		callError = NewMessageCallError(messageCall.UniqueID, ErrorCodeNOTSUPPORTED, "", NoError{})
	}

	s.sendResponse(packet.Client, callResult, callError)
}

func (s *Ocpp) sendResponse(client *ws.Client, callResult *MessageCallResult, callError *MessageCallError) {
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
