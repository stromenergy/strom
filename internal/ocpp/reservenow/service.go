package reservenow

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/call"
)

type ReserveNow struct {
	repository *db.Repository
	call       *call.Call
}

func NewService(repository *db.Repository, callService *call.Call) *ReserveNow {
	return &ReserveNow{
		repository: repository,
		call:       callService,
	}
}
