package ws

import (
	"bytes"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

const (
	WRITE_WAIT = 10 * time.Second
	PONG_WAIT = 60 * time.Second
	PING_INTERVAL = (PONG_WAIT * 9) / 10
	READ_LIMIT = 512
)

var (
	NEWLINE = []byte{'\n'}
	SPACE   = []byte{' '}
)

type Client struct {
	Send       chan []byte
	conn       *websocket.Conn
	dispatcher Dispatcher
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
				log.Error().Msg("STR020: Error unexpected websocket close")
				log.Error().Err(err)
			}

			break
		}
		
		message = bytes.TrimSpace(bytes.Replace(message, NEWLINE, SPACE, -1))
		c.dispatcher.Broadcast(message)
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
		case message, ok := <-c.Send:
			c.conn.SetWriteDeadline(time.Now().Add(WRITE_WAIT))

			if !ok {
				// The dispatcher has closed the channel
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				log.Error().Msg("STR021: Error getting writer for the message")
				log.Error().Err(err)
				return
			}

			w.Write(message)

			// Add all the other queued messages to the writer
			n := len(c.Send)

			for i := 0; i < n; i++ {
				w.Write(NEWLINE)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				log.Error().Msg("STR022: Error closing the writer")
				log.Error().Err(err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(WRITE_WAIT))

			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Error().Msg("STR021: Error writing ping message")
				log.Error().Err(err)
				return
			}
		}
	}
}
