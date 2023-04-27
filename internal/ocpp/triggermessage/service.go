package triggermessage

import "github.com/stromenergy/strom/internal/db"

type TriggerMessage struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *TriggerMessage {
	return &TriggerMessage{
		repository: repository,
	}
}
