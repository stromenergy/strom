package datatransfer

import "github.com/stromenergy/strom/internal/db"

type DataTransfer struct {
	repository *db.Repository
	handlers   map[string]Handler
}

func NewService(repository *db.Repository) *DataTransfer {
	return &DataTransfer{
		repository: repository,
		handlers:   make(map[string]Handler),
	}
}
