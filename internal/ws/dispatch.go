package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Dispatcher interface {
	CheckOrigin(r *http.Request) bool
	Packet(packet *Packet)
	Register(client *Client, params gin.Params)
	Subprotocols() []string
	Unregister(client *Client)
}