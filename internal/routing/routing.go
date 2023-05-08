package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/stromenergy/strom/frontend"
)

func (s *Routing) mountRoutes(engine *gin.Engine) {
	s.services.MountRoutes(engine)
}

func (s *Routing) mountStaticRoutes(engine *gin.Engine) {
	frontend.MountRoutes(engine)
}
