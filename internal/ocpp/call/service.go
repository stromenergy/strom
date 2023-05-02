package call

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
)

type Call struct {
	repository *db.Repository
	channels   map[string]chan types.Message
}

func NewService(repository *db.Repository) *Call {
	return &Call{
		repository: repository,
		channels:   make(map[string]chan types.Message),
	}
}
