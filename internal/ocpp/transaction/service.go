package transaction

import (
	"github.com/stromenergy/strom/internal/db"
)

type Transaction struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *Transaction {
	return &Transaction{
		repository: repository,
	}
}
