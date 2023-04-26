package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stromenergy/strom/internal/util"
)

func WebsocketHandler(ctx *gin.Context, dispatcher Dispatcher) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		Subprotocols: dispatcher.Subprotocols(),
		CheckOrigin: dispatcher.CheckOrigin,
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		util.LogError("STR018: Error upgrading websocket", err)
		return
	}

	client := &Client{
		Send:       make(chan []byte, 256),
		conn:       conn,
		dispatcher: dispatcher,
	}

	dispatcher.Register(client, ctx.Params)

	go client.reader()
	go client.writer()
}
