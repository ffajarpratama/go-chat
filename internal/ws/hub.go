package ws

import (
	"github.com/google/uuid"
)

type Hub struct {
	Chatrooms     map[uuid.UUID]*Chatroom
	Register      chan *Client
	Unregister    chan *Client
	BroadcastChan chan *Message
}

type Chatroom struct {
	ChatroomID uuid.UUID
	Clients    map[uuid.UUID]*Client

	// Sender     *Client
	// Receiver   *Client
}

type Message struct {
	ChatroomID uuid.UUID `json:"chatroom_id"`
	SenderID   uuid.UUID `json:"sender_id"`
	ReceiverID uuid.UUID `json:"receiver_id"`
	Content    string    `json:"content"`
}

func NewHub() *Hub {
	return &Hub{
		Chatrooms:     make(map[uuid.UUID]*Chatroom),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		BroadcastChan: make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if room, roomOK := h.Chatrooms[client.ChatroomID]; roomOK {
				if _, clientOK := room.Clients[client.ID]; !clientOK {
					room.Clients[client.ID] = client
				}
			}
		case client := <-h.Unregister:
			if _, roomOK := h.Chatrooms[client.ChatroomID]; roomOK {
				if _, clientOK := h.Chatrooms[client.ChatroomID].Clients[client.ID]; clientOK {
					delete(h.Chatrooms[client.ChatroomID].Clients, client.ID)
					close(client.MsgChan)
				}
			}
		case msg := <-h.BroadcastChan:
			room, ok := h.Chatrooms[msg.ChatroomID]
			if ok {
				room.Clients[msg.ReceiverID].MsgChan <- msg
				// for _, v := range room.Clients {
				// 	v.MsgChan <- msg
				// }
			}
		}
	}
}
