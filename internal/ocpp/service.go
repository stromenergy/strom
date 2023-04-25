package ocpp

import (
	"context"
	"sync"

	"github.com/stromenergy/strom/internal/db"
)

type Ocpp interface {
	Start(context.Context, *sync.WaitGroup)
}

type OcppService struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) Ocpp {
	return &OcppService{
		repository: repository,
	}
}

func (s *OcppService) Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
}
