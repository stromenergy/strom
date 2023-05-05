package management

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/call"
)

type Management struct {
	repository *db.Repository
	call       *call.Call
}

func NewService(repository *db.Repository, callService *call.Call) *Management {
	return &Management{
		repository: repository,
		call:       callService,
	}
}
