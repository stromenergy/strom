package diagnostic

import (
	"context"

	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Diagnostic) DiagnosticsStatusNotificationReq(client *ws.Client, message types.Message) {
	_, err := UnmarshalDiagnosticsStatusNotificationReq(message.Payload)

	if err != nil {
		util.LogError("STR058: Error unmarshaling DiagnosticsStatusNotificationReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	_, err = s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR059: Charge point not found", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	// TODO: Do something with the diagnostics status
	callResult := types.NewMessageCallResult(message.UniqueID, DiagnosticsStatusNotificationConf{})
	callResult.Send(client)
}
