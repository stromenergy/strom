package authorization

import (
	"github.com/stromenergy/strom/internal/db"
)

type Authorization struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *Authorization {
	return &Authorization{
		repository: repository,
	}
}
