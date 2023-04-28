package rest

import (
	"github.com/gin-gonic/gin"
)

func (s *Rest) mountRoutes(engine *gin.Engine) {
	s.services.MountRoutes(engine)
}
