package reservation

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/call"
)

type Reservation struct {
	repository *db.Repository
	call       *call.Call
}

func NewService(repository *db.Repository, callService *call.Call) *Reservation {
	return &Reservation{
		repository: repository,
		call:       callService,
	}
}
