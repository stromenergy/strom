package heartbeat

import "github.com/stromenergy/strom/internal/db"

type Heartbeat struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *Heartbeat {
	return &Heartbeat{
		repository: repository,
	}
}
