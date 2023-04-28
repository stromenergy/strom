package call

import (
	"github.com/stromenergy/strom/internal/db"
)

type Call struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *Call {
	return &Call{
		repository: repository,
	}
}
