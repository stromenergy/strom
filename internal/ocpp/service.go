package ocpp

import (
	"context"
	"sync"

	"github.com/stromenergy/strom/internal/db"
)

type OcppInterface interface {
	Start(shutdownCtx context.Context,waitGroup  *sync.WaitGroup)
}

type Ocpp struct {
	repository *db.Repository
}

func NewService(repository *db.Repository) OcppInterface {
	return &Ocpp{
		repository: repository,
	}
}

func (s *Ocpp) Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
}
