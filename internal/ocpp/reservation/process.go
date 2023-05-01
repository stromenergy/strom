package reservation

import (
	"context"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Reservation) CancelReservationConf(client *ws.Client, message types.Message) {
	cancelReservationConf, err := unmarshalCancelReservationConf(message.Payload)

	if err != nil {
		util.LogError("STR064: Error unmarshaling CancelReservationConf", err)
		return
	}

	ctx := context.Background()
	reservation, err := s.repository.GetReservationByReqID(ctx, message.UniqueID)

	if err != nil {
		util.LogError("STR065: Error getting reservation", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	// TODO: Handle conf response with a channel
	if cancelReservationConf.Status == types.CancelReservationStatusAccepted {
		updateReservationParams := param.NewUpdateReservationParams(reservation)
		updateReservationParams.Status = db.ReservationStatusCancelled
	
		_, err = s.repository.UpdateReservation(ctx, updateReservationParams)
	
		if err != nil {
			util.LogError("STR066: Error updating reservation", err)
			return
		}
	}
}

func (s *Reservation) ReserveNowConf(client *ws.Client, message types.Message) {
	reserveNowConf, err := unmarshalReserveNowConf(message.Payload)

	if err != nil {
		util.LogError("STR042: Error unmarshaling ReserveNowConf", err)
		return
	}

	ctx := context.Background()
	reservation, err := s.repository.GetReservationByReqID(ctx, message.UniqueID)

	if err != nil {
		util.LogError("STR043: Error getting reservation", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	updateReservationParams := param.NewUpdateReservationParams(reservation)
	updateReservationParams.Status = db.ReservationStatus(reserveNowConf.Status)

	_, err = s.repository.UpdateReservation(ctx, updateReservationParams)

	if err != nil {
		util.LogError("STR044: Error updating reservation", err)
		return
	}
}
