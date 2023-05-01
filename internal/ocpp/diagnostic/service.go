package diagnostic

import (
	"github.com/stromenergy/strom/internal/db"
)

type Diagnostic struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *Diagnostic {
	return &Diagnostic{
		repository: repository,
	}
}
