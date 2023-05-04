package firmware

import (
	"context"

	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Firmware) FirmwareStatusNotificationReq(client *ws.Client, message types.Message) {
	_, err := unmarshalFirmwareStatusNotificationReq(message.Payload)

	if err != nil {
		util.LogError("STR060: Error unmarshaling FirmwareStatusNotificationReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	_, err = s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR061: Charge point not found", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	// TODO: Do something with the firmware status
	callResult := types.NewMessageCallResult(message.UniqueID, FirmwareStatusNotificationConf{})
	callResult.Send(client)
}
