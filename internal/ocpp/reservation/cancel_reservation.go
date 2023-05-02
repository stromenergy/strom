package reservation

import (
	"context"

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

	cancelReservationReq := CancelReservationReq{
		ReservationID: reservation.ID,
	}

	uniqueID, channel := s.call.Send(client, types.CallActionCancelReservation, cancelReservationReq)

	// Wait for the channel to produce a response
	go s.waitForCancelReservationConf(client, reservation, uniqueID, channel)
}

func (s *Reservation) waitForCancelReservationConf(client *ws.Client, reservation db.Reservation, uniqueID string, channel <-chan types.Message) {
	message := <-channel
	
	// Remove the channel as a response listener
	s.call.Remove(uniqueID)

	// Update the reservation
	ctx := context.Background()
	updateReservationParams := param.NewUpdateReservationParams(reservation)

	if message.MessageType == types.MessageTypeCallResult {
		cancelReservationConf, err := unmarshalCancelReservationConf(message.Payload)

		if err != nil {
			util.LogError("STR064: Error unmarshaling CancelReservationConf", err)
			return
		}
	
		updateReservationParams.Status = db.ReservationStatus(cancelReservationConf.Status)

		_, err = s.repository.UpdateReservation(ctx, updateReservationParams)

		if err != nil {
			util.LogError("STR044: Error updating reservation", err)
		}
	
	}
}
