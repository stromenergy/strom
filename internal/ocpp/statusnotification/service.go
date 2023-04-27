package statusnotification

import "github.com/stromenergy/strom/internal/db"

type StatusNotification struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *StatusNotification {
	return &StatusNotification{
		repository: repository,
	}
}
