package metervalue

import (
	"github.com/stromenergy/strom/internal/db"
)

type MeterValue struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) *MeterValue {
	return &MeterValue{
		repository: repository,
	}
}
