package triggermessage

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/call"
)

type TriggerMessage struct {
	repository *db.Repository
	call       *call.Call
}

func NewService(repository *db.Repository, callService *call.Call) *TriggerMessage {
	return &TriggerMessage{
		repository: repository,
		call:       callService,
	}
}
