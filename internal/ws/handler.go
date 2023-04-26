package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

func WebsocketHandler(ctx *gin.Context, dispatcher Dispatcher) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Error().Msg("STR018: Error upgrading websocket")
		log.Error().Err(err)
		return
	}

	client := &Client{
		Send:       make(chan []byte, 256),
		conn:       conn,
		dispatcher: dispatcher,
	}

	dispatcher.Register(client)

	go client.reader()
	go client.writer()
}
