package transaction

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/metervalue"
)

type Transaction struct {
	repository *db.Repository
	meterValue *metervalue.MeterValue
}

func NewService(repository *db.Repository, meterValue *metervalue.MeterValue) *Transaction {
	return &Transaction{
		repository: repository,
		meterValue: meterValue,
	}
}
