package bootnotification

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/triggermessage"
)

type BootNotification struct {
	repository     *db.Repository
	triggerMessage *triggermessage.TriggerMessage
}

func NewService(repository *db.Repository, triggerMessage *triggermessage.TriggerMessage) *BootNotification {
	return &BootNotification{
		repository:     repository,
		triggerMessage: triggerMessage,
	}
}
