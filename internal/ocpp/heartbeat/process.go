package heartbeat

import (
	"context"

	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Heartbeat) ProcessReq(client *ws.Client, message types.Message) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err == nil {
		// Update existing charge point
		updateChargePointParams := param.NewUpdateChargePointParams(chargePoint)

		_, err = s.repository.UpdateChargePoint(ctx, updateChargePointParams)

		if err != nil {
			util.LogError("5: Error updating charge point", err)
			callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
			callError.Send(client)
			return
		}
	}

	// TODO: Notify UI of changes

	heartbeatConf := HeartbeatConf{
		CurrentTime: types.NewOcppTime(nil),
	}

	callResult := types.NewMessageCallResult(message.UniqueID, heartbeatConf)
	callResult.Send(client)
}
