package ocpp

import (
	"encoding/json"
	"fmt"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Ocpp) processPacket(packet *ws.Packet) {
	message := types.Message{}

	if err := json.Unmarshal(packet.Message, &message); err != nil {
		util.LogError("STR030: Error unmarshaling message", err)
		return
	}

	switch message.MessageType {
	case types.MessageTypeCall:
		s.processCall(packet.Client, message)
	case types.MessageTypeCallError:
		s.processCallError(packet.Client, message)
	case types.MessageTypeCallResult:
		s.processCallResult(packet.Client, message)
	}
}

func (s *Ocpp) processCall(client *ws.Client, message types.Message) {
	switch message.Action {
	case db.CallActionBootNotification:
		s.bootNotification.BootNotificationReq(client, message)
	case db.CallActionHeartbeat:
		s.heartbeat.HeartbeatReq(client, message)
	case db.CallActionMeterValues:
		s.meterValue.MeterValueReq(client, message)
	case db.CallActionStartTransaction:
		s.transaction.StartTransactionReq(client, message)
	case db.CallActionStatusNotification:
		s.statusNotification.StatusNotificationReq(client, message)
	case db.CallActionStopTransaction:
		s.transaction.StopTransactionReq(client, message)
	default:
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeNotSupported, "", nil)
		callError.Send(client)
	}
}

func (s *Ocpp) processCallError(client *ws.Client, message types.Message) {
	if _, err := s.call.Find(client, message); err != nil {
		// TODO: Handle error reponses per action
		util.LogError(fmt.Sprintf("STR031: Error response %s, %s, %s", message.UniqueID, message.ErrorCode, message.ErrorDescription), nil)
	}
}

func (s *Ocpp) processCallResult(client *ws.Client, message types.Message) {
	if call, err := s.call.Find(client, message); err != nil {
		switch call.Action {
		case db.CallActionReserveNow:
			s.reserverNow.ReserveNowConf(client, message)
		case db.CallActionTriggerMessage:
			s.triggerMessage.TriggerMessageConf(client, message)
		}
	}
}
