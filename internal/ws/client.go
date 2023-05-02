package ws

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stromenergy/strom/internal/util"
)

const (
	WRITE_WAIT    = 10 * time.Second
	PONG_WAIT     = 60 * time.Second
	PING_INTERVAL = (PONG_WAIT * 9) / 10
	READ_LIMIT    = 512
)

var (
	NEWLINE = []byte{'\n'}
	SPACE   = []byte{' '}
)

type Client struct {
	ID         string
	conn       *websocket.Conn
	dispatcher Dispatcher
	queue      chan []byte
}

func (c *Client) CloseQueue() {
	close(c.queue)
}

func (c *Client) Send(message []byte) {
	c.queue <- message
}

func (c *Client) close() {
	c.dispatcher.Unregister(c)
	c.conn.Close()
}

func (c *Client) pongHandler(appData string) error {
	c.conn.SetReadDeadline(time.Now().Add(PONG_WAIT))
	return nil
}

func (c *Client) reader() {
	defer c.close()

	c.conn.SetPongHandler(c.pongHandler)
	c.conn.SetReadDeadline(time.Now().Add(PONG_WAIT))
	c.conn.SetReadLimit(READ_LIMIT)

	for {
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				util.LogError("STR020: Error unexpected websocket close", err)
			}

			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, NEWLINE, SPACE, -1))
		util.LogDebug(fmt.Sprintf("> %s", string(message)))
		c.dispatcher.Packet(&Packet{Client: c, Message: message})
	}
}

func (c *Client) writer() {
	ticker := time.NewTicker(PING_INTERVAL)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.queue:
			c.conn.SetWriteDeadline(time.Now().Add(WRITE_WAIT))

			if !ok {
				// The dispatcher has closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				util.LogError("STR021: Error getting writer for the message", err)
				return
			}

			util.LogDebug(fmt.Sprintf("< %s", string(message)))
			w.Write(message)

			if err := w.Close(); err != nil {
				util.LogError("STR022: Error closing the writer", err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(WRITE_WAIT))

			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				util.LogError("STR031: Error writing ping message", err)
				return
			}
		}
	}
}
