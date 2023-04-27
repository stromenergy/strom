package heartbeat

import (
	"context"

	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
)

func (s *Heartbeat) ProcessReq(clientID string, messageCall types.MessageCall) (*types.MessageCallResult, *types.MessageCallError) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, clientID)

	if err == nil {
		// Update existing charge point
		updateChargePointParams := param.NewUpdateChargePointParams(chargePoint)

		_, err = s.repository.UpdateChargePoint(ctx, updateChargePointParams)

		if err != nil {
			util.LogError("STR033: Error updating charge point", err)
			return nil, types.NewMessageCallError(messageCall.UniqueID, types.ErrorCodeINTERNALERROR, "", types.NoError{})
		}
	}

	// TODO: Notify UI of changes

	heartbeatConf := HeartbeatConf{
		CurrentTime: types.NewOcppTime(nil),
	}

	return types.NewMessageCallResult(messageCall.UniqueID, heartbeatConf), nil
}
