package rest

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/service"
)

type RestInterface interface {
	Start(port string, shutdownCtx context.Context, waitGroup *sync.WaitGroup)
}

type Rest struct {
	repository  *db.Repository
	server      *http.Server
	services    service.ServiceInterface
	shutdownCtx context.Context
	waitGroup   *sync.WaitGroup
}

func NewService(repository *db.Repository, services service.ServiceInterface) RestInterface {
	return &Rest{
		repository: repository,
		services:   services,
	}
}

func (s *Rest) Start(port string, shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	s.shutdownCtx = shutdownCtx
	s.waitGroup = waitGroup
	waitGroup.Add(1)

	go s.listenAndServe(port)
	go s.waitForShutdown()
}

func (s *Rest) listenAndServe(port string) {
	engine := gin.Default()

	s.mountRoutes(engine)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	err := s.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Error().Msg("STR018: Shutting down Rest service")
	}
}

func (s *Rest) waitForShutdown() {
	<-s.shutdownCtx.Done()
	log.Debug().Msg("Shutting down Rest service")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)

	if err != nil {
		log.Error().Msg("STR019: Shutting down Rest service")
	}

	log.Printf("Rest service shut down")
	s.waitGroup.Done()
}
