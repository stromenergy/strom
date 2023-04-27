package ocpp

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/bootnotification"
	"github.com/stromenergy/strom/internal/ocpp/heartbeat"
	"github.com/stromenergy/strom/internal/ocpp/triggermessage"
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
	inbox      chan *ws.Packet
	clients    map[*ws.Client]bool
	register   chan *ws.Client
	unregister chan *ws.Client
	// Services
	bootNotification *bootnotification.BootNotification
	heartbeat        *heartbeat.Heartbeat
	triggerMessage   *triggermessage.TriggerMessage
}

func NewService(repository *db.Repository) OcppInterface {
	return &Ocpp{
		repository: repository,
		// Dispatcher
		inbox:      make(chan *ws.Packet),
		register:   make(chan *ws.Client),
		unregister: make(chan *ws.Client),
		clients:    make(map[*ws.Client]bool),
		// Services
		bootNotification: bootnotification.NewService(repository),
		heartbeat:        heartbeat.NewService(repository),
		triggerMessage:   triggermessage.NewService(repository),
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
