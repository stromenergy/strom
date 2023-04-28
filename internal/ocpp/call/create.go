package call

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/util"
)

func (s *Call) Create(chargePointId int64, action db.CallAction) (db.Call, error) {
	ctx := context.Background()
	createCallParams := db.CreateCallParams{
		ReqID:         uuid.NewString(),
		ChargePointID: chargePointId,
		Action:        action,
		CreatedAt:     time.Now(),
	}

	call, err := s.repository.CreateCall(ctx, createCallParams)

	if err != nil {
		util.LogError("STR045: Error creating call", err)
	}

	return call, err
}
