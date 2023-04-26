package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Dispatcher interface {
	Broadcast(message []byte)
	CheckOrigin(r *http.Request) bool
	Register(client *Client, params gin.Params)
	Subprotocols() []string
	Unregister(client *Client)
}