package ocpp

import (
	"github.com/gin-gonic/gin"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Ocpp) WebsocketHandler(ctx *gin.Context) {
	ws.WebsocketHandler(ctx, s)
}
