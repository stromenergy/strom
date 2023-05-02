package datatransfer

import (
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/call"
)

type DataTransfer struct {
	repository *db.Repository
	call       *call.Call
	handlers   map[string]Handler
}

func NewService(repository *db.Repository, callService *call.Call) *DataTransfer {
	return &DataTransfer{
		repository: repository,
		call:       callService,
		handlers:   make(map[string]Handler),
	}
}
