package transaction

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/authorization"
	"github.com/stromenergy/strom/internal/ocpp/metervalue"
)

type Transaction struct {
	repository    *db.Repository
	authorization *authorization.Authorization
	meterValue    *metervalue.MeterValue
}

func NewService(repository *db.Repository, authorization *authorization.Authorization, meterValue *metervalue.MeterValue) *Transaction {
	return &Transaction{
		repository: repository,
		authorization: authorization,
		meterValue: meterValue,
	}
}
