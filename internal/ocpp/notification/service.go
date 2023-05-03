package notification

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/call"
	"github.com/stromenergy/strom/internal/ocpp/triggermessage"
)

type Notification struct {
	repository     *db.Repository
	call           *call.Call
	triggerMessage *triggermessage.TriggerMessage
}

func NewService(repository *db.Repository, callService *call.Call, triggerMessage *triggermessage.TriggerMessage) *Notification {
	return &Notification{
		repository:     repository,
		call:           callService,
		triggerMessage: triggerMessage,
	}
}
