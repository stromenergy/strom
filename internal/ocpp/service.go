package ocpp

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ws"
)

type OcppInterface interface {
	Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup)
	MountRoutes(engine *gin.Engine)
}

type Ocpp struct {
	repository  *db.Repository
	shutdownCtx context.Context
	waitGroup   *sync.WaitGroup

	// Dispatcher
	broadcast  chan []byte
	clients    map[*ws.Client]bool
	register   chan *ws.Client
	unregister chan *ws.Client
}

func NewService(repository *db.Repository) OcppInterface {
	return &Ocpp{
		repository: repository,
		// Dispatcher
		broadcast:  make(chan []byte),
		register:   make(chan *ws.Client),
		unregister: make(chan *ws.Client),
		clients:    make(map[*ws.Client]bool),
	}
}

func (s *Ocpp) Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	s.shutdownCtx = shutdownCtx
	s.waitGroup = waitGroup
	waitGroup.Add(1)

	go s.dispatch()
	go s.waitForShutdown()
}

func (s *Ocpp) waitForShutdown() {
	<-s.shutdownCtx.Done()
	log.Debug().Msg("Shutting down Ocpp service")

	for client := range s.clients {
		s.remove(client)
	}

	log.Printf("Ocpp service shut down")
	s.waitGroup.Done()
}
