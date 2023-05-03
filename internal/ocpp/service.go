package ocpp

import (
	"context"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/authorization"
	"github.com/stromenergy/strom/internal/ocpp/call"
	"github.com/stromenergy/strom/internal/ocpp/datatransfer"
	"github.com/stromenergy/strom/internal/ocpp/diagnostic"
	"github.com/stromenergy/strom/internal/ocpp/firmware"
	"github.com/stromenergy/strom/internal/ocpp/heartbeat"
	"github.com/stromenergy/strom/internal/ocpp/metervalue"
	"github.com/stromenergy/strom/internal/ocpp/notification"
	"github.com/stromenergy/strom/internal/ocpp/reservation"
	"github.com/stromenergy/strom/internal/ocpp/transaction"
	"github.com/stromenergy/strom/internal/ocpp/triggermessage"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

type OcppInterface interface {
	Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup)
	MountRoutes(engine *gin.Engine)

	CancelReservation(client *ws.Client, reservationID int64)
	ChangeAvailability(client *ws.Client, connectorId int32, availabilityType types.AvailabilityType) (string, <-chan types.Message)
	DataTransfer(client *ws.Client, vendorId string, messageID, data *string) (string, <-chan types.Message)
	ReserveNow(client *ws.Client, connectorId int32, expiryDate time.Time, idTag string, parentIdTag *string)
	TriggerMessage(client *ws.Client, chargePointId int64, messageTrigger types.MessageTrigger, connectorId *int32) (string, <-chan types.Message)
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
	call           *call.Call
	authorization  *authorization.Authorization
	dataTransfer   *datatransfer.DataTransfer
	diagnostic     *diagnostic.Diagnostic
	firmware       *firmware.Firmware
	heartbeat      *heartbeat.Heartbeat
	meterValue     *metervalue.MeterValue
	notification   *notification.Notification
	reservation    *reservation.Reservation
	transaction    *transaction.Transaction
	triggerMessage *triggermessage.TriggerMessage
}

func NewService(repository *db.Repository) OcppInterface {
	authorization := authorization.NewService(repository)
	callService := call.NewService(repository)
	meterValue := metervalue.NewService(repository)
	triggerMessageService := triggermessage.NewService(repository, callService)

	return &Ocpp{
		repository: repository,
		// Dispatcher
		inbox:      make(chan *ws.Packet),
		register:   make(chan *ws.Client),
		unregister: make(chan *ws.Client),
		clients:    make(map[*ws.Client]bool),
		// Services
		authorization:  authorization,
		call:           callService,
		dataTransfer:   datatransfer.NewService(repository, callService),
		diagnostic:     diagnostic.NewService(repository),
		firmware:       firmware.NewService(repository),
		heartbeat:      heartbeat.NewService(repository),
		meterValue:     meterValue,
		notification:   notification.NewService(repository, callService, triggerMessageService),
		reservation:    reservation.NewService(repository, callService),
		transaction:    transaction.NewService(repository, authorization, meterValue),
		triggerMessage: triggerMessageService,
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
