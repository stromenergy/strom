package reservation

import (
	"context"
	"errors"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/call"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Reservation) SendCancelReservationReq(client *ws.Client, reservationID int64) (string, <-chan types.Message, error) {
	ctx := context.Background()
	reservation, err := s.repository.GetReservation(ctx, reservationID)

	if err != nil {
		util.LogError("STR062: Reservation not found", err)
		return "", nil, errors.New("Reservation not found")
	}

	cancelReservationReq := CancelReservationReq{
		ReservationID: reservation.ID,
	}

	uniqueID, confChannel, err := s.call.Send(client, types.CallActionCancelReservation, cancelReservationReq)

	if err != nil {
		s.call.Remove(uniqueID)
		return "", nil, err
	}

	// Wait for the channel to produce a response
	reqChan := make(chan types.Message)

	go s.waitForCancelReservationConf(client, reservation, uniqueID, confChannel, reqChan)

	return uniqueID, reqChan, nil
}

func (s *Reservation) waitForCancelReservationConf(client *ws.Client, reservation db.Reservation, uniqueID string, confChannel <-chan types.Message, reqChannel chan<- types.Message) {
	message := <-confChannel

	// Forward message to requestor
	defer call.Forward(message, reqChannel)

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
