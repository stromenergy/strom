package ocpp

import "github.com/gin-gonic/gin"

func (s *Ocpp) MountRoutes(engine *gin.Engine) {
	ocpp := engine.Group("/ocpp")
	{
		ocpp.GET("/ws", s.WebsocketHandler)
	}
}