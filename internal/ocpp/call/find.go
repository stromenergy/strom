package call

import (
	"context"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Call) Find(client *ws.Client, message types.Message) (db.Call, error) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		return db.Call{}, err
	}

	getCallByReqIDParams := db.GetCallByReqIDParams{
		ChargePointID: chargePoint.ID,
		ReqID:         message.UniqueID,
	}

	return s.repository.GetCallByReqID(ctx, getCallByReqIDParams)
}
