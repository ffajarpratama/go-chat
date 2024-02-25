package ws

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID         uuid.UUID
	ChatroomID uuid.UUID
	Conn       *websocket.Conn
	MsgChan    chan *Message
	Chatroom   *Chatroom
}

func (c *Client) Write() {
	defer c.Conn.Close()

	for {
		msg, ok := <-c.MsgChan
		if !ok {
			return
		}

		c.Conn.WriteJSON(msg)
	}
}

func (c *Client) Read(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg *Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("err: %v\n", err)
			}

			break
		}

		hub.BroadcastChan <- msg
	}
}
