package service

import (
	"context"
	"sync"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp"
)

type ServiceInterface interface {
	Start(shutdownCtx context.Context,waitGroup  *sync.WaitGroup)
}

type Service struct {
	Ocpp ocpp.OcppInterface
}

func NewService(repository *db.Repository) ServiceInterface {
	ocppService := ocpp.NewService(repository)

	return &Service{
		Ocpp: ocppService,
	}
}

func (s *Service) Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	s.Ocpp.Start(shutdownCtx, waitGroup)
}