package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Dispatcher interface {
	CheckOrigin(r *http.Request) bool
	Message(message []byte)
	Register(client *Client, params gin.Params)
	Subprotocols() []string
	Unregister(client *Client)
}