package service

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp"
)

type ServiceInterface interface {
	Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup)
	MountRoutes(engine *gin.Engine)
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

func (s *Service) MountRoutes(engine *gin.Engine) {
	s.Ocpp.MountRoutes(engine)
}
