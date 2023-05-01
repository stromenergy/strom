package notification

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/triggermessage"
)

type Notification struct {
	repository     *db.Repository
	triggerMessage *triggermessage.TriggerMessage
}

func NewService(repository *db.Repository, triggerMessage *triggermessage.TriggerMessage) *Notification {
	return &Notification{
		repository:     repository,
		triggerMessage: triggerMessage,
	}
}
