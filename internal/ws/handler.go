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
		Subprotocols:    dispatcher.Subprotocols(),
		CheckOrigin:     dispatcher.CheckOrigin,
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		util.LogError("STR001: Error upgrading websocket", err)
		return
	}

	client := &Client{
		conn:       conn,
		dispatcher: dispatcher,
		queue:      make(chan []byte, 256),
	}

	dispatcher.Register(client, ctx.Params)

	go client.reader()
	go client.writer()
}
