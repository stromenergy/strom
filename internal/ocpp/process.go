package ocpp

import (
	"encoding/json"

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
	default:
		s.call.Receive(packet.Client, message)
	}
}

func (s *Ocpp) processCall(client *ws.Client, message types.Message) {
	switch message.Action {
	case types.CallActionAuthorize:
		s.authorization.AuthorizeReq(client, message)
	case types.CallActionBootNotification:
		s.notification.BootNotificationReq(client, message)
	case types.CallActionDataTransfer:
		s.dataTransfer.DataTransferReq(client, message)
	case types.CallActionDiagnosticsStatusNotification:
		s.diagnostic.DiagnosticsStatusNotificationReq(client, message)
	case types.CallActionFirmwareStatusNotification:
		s.firmware.FirmwareStatusNotificationReq(client, message)
	case types.CallActionHeartbeat:
		s.heartbeat.HeartbeatReq(client, message)
	case types.CallActionMeterValues:
		s.meterValue.MeterValueReq(client, message)
	case types.CallActionStartTransaction:
		s.transaction.StartTransactionReq(client, message)
	case types.CallActionStatusNotification:
		s.notification.StatusNotificationReq(client, message)
	case types.CallActionStopTransaction:
		s.transaction.StopTransactionReq(client, message)
	default:
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeNotSupported, "", nil)
		callError.Send(client)
	}
}
