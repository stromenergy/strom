package bootnotification

import "github.com/stromenergy/strom/internal/db"


type BootNotification struct {
	repository  *db.Repository
}

func NewService(repository *db.Repository) *BootNotification {
	return &BootNotification{
		repository: repository,
	}
}