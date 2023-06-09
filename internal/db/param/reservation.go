package param

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
)

func NewUpdateReservationParams(reservation db.Reservation) db.UpdateReservationParams {
	return db.UpdateReservationParams{
		ID:        reservation.ID,
		Status:    reservation.Status,
		UpdatedAt: time.Now(),
	}
}
