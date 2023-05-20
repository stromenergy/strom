package routing

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
	"github.com/stromenergy/strom/internal/util"
)

type RoutingInterface interface {
	Start(port string, tlsCertificateFile, tlsPrivateKeyFile *string, shutdownCtx context.Context, waitGroup *sync.WaitGroup)
}

type Routing struct {
	repository  *db.Repository
	server      *http.Server
	services    service.ServiceInterface
	shutdownCtx context.Context
	waitGroup   *sync.WaitGroup
}

func NewService(repository *db.Repository, services service.ServiceInterface) RoutingInterface {
	return &Routing{
		repository: repository,
		services:   services,
	}
}

func (s *Routing) Start(port string, tlsCertificateFile, tlsPrivateKeyFile *string, shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	s.shutdownCtx = shutdownCtx
	s.waitGroup = waitGroup
	waitGroup.Add(1)

	go s.listenAndServe(port, tlsCertificateFile, tlsPrivateKeyFile)
	go s.waitForShutdown()
}

func (s *Routing) listenAndServe(port string, tlsCertificateFile, tlsPrivateKeyFile *string) {
	engine := gin.Default()

	s.mountRoutes(engine)
	s.mountStaticRoutes(engine)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	var err error

	if tlsCertificateFile != nil && tlsPrivateKeyFile != nil {
		err = s.server.ListenAndServeTLS(*tlsCertificateFile, *tlsPrivateKeyFile)
	} else {
		err = s.server.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		util.LogError("STR018: Shutting down Rest service", err)
	}
}

func (s *Routing) waitForShutdown() {
	<-s.shutdownCtx.Done()
	log.Debug().Msg("Shutting down Rest service")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)

	if err != nil {
		util.LogError("STR019: Shutting down Rest service", err)
	}

	log.Printf("Rest service shut down")
	s.waitGroup.Done()
}
