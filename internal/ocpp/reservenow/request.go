package reservenow

import (
	"context"
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *ReserveNow) Request(client *ws.Client, connectorId int32, expiryDate time.Time, idTag string, parentIdTag *string) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR040: Charge point not found", err)
		return
	}

	if call, err := s.call.Create(chargePoint.ID, db.CallActionTriggerMessage); err == nil {
		createReservationParams := db.CreateReservationParams{
			ChargePointID: chargePoint.ID,
			ConnectorID:   connectorId,
			ReqID:         call.ReqID,
			ExpiryDate:    expiryDate,
			Status:        db.ReservationStatusAccepted,
			IDTag:         idTag,
			ParentIDTag:   util.SqlNullString(parentIdTag),
		}

		reservation, err := s.repository.CreateReservation(ctx, createReservationParams)

		if err != nil {
			util.LogError("STR041: Error creating reservation", err)
			return
		}

		reserveNowReq := ReserveNowReq{
			ConnectorID:   connectorId,
			ExpiryDate:    types.NewOcppTime(&expiryDate),
			IDTag:         idTag,
			ParentIDTag:   parentIdTag,
			ReservationID: reservation.ID,
		}

		message := types.NewMessageCall(call.ReqID, db.CallActionReserveNow, reserveNowReq)
		message.Send(client)
	}
}
