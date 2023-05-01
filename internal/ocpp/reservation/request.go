package reservation

import (
	"context"
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Reservation) SendCancelReservationReq(client *ws.Client, reservationID int64) {
	ctx := context.Background()
	reservation, err := s.repository.GetReservation(ctx, reservationID)

	if err != nil {
		util.LogError("STR062: reservation not found", err)
		// TODO: Return an error
		return
	}

	if call, err := s.call.Create(reservation.ChargePointID, db.CallActionTriggerMessage); err == nil {
		updateReservationParams := param.NewUpdateReservationParams(reservation)
		updateReservationParams.ReqID = call.ReqID

		_, err := s.repository.UpdateReservation(ctx, updateReservationParams)

		if err != nil {
			util.LogError("STR063: Error updating reservation", err)
			return
		}

		cancelReservationReq := CancelReservationReq{
			ReservationID: reservation.ID,
		}

		message := types.NewMessageCall(call.ReqID, db.CallActionCancelReservation, cancelReservationReq)
		message.Send(client)
	}
}

func (s *Reservation) SendReserveNowReq(client *ws.Client, connectorId int32, expiryDate time.Time, idTag string, parentIdTag *string) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR040: Charge point not found", err)
		return
	}

	// TODO: check IDTag can make reservation
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
