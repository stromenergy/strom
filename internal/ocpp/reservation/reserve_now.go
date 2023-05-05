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

func (s *Reservation) SendReserveNowReq(client *ws.Client, connectorId int32, expiryDate time.Time, idTag string, parentIdTag *string) {
	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR040: Charge point not found", err)
		return
	}

	createReservationParams := db.CreateReservationParams{
		ChargePointID: chargePoint.ID,
		ConnectorID:   connectorId,
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

	uniqueID, channel := s.call.Send(client, types.CallActionReserveNow, reserveNowReq)

	// Wait for the channel to produce a response
	go s.waitForReserveNowConf(client, reservation, uniqueID, channel)
}

func (s *Reservation) waitForReserveNowConf(client *ws.Client, reservation db.Reservation, uniqueID string, channel <-chan types.Message) {
	message := <-channel

	// Update the reservation
	ctx := context.Background()
	updateReservationParams := param.NewUpdateReservationParams(reservation)

	switch message.MessageType {
	case types.MessageTypeCallError:
		updateReservationParams.Status = db.ReservationStatusRejected
	case types.MessageTypeCallResult:
		reserveNowConf, err := unmarshalReserveNowConf(message.Payload)

		if err != nil {
			util.LogError("STR042: Error unmarshaling ReserveNowConf", err)
			return
		}

		updateReservationParams.Status = db.ReservationStatus(reserveNowConf.Status)
	}

	_, err := s.repository.UpdateReservation(ctx, updateReservationParams)

	if err != nil {
		util.LogError("STR043: Error updating reservation", err)
	}
}
