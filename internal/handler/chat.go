package handler

import (
	"net/http"

	"github.com/ffajarpratama/go-chat/internal/ws"
	custom_error "github.com/ffajarpratama/go-chat/pkg/error"
	responser "github.com/ffajarpratama/go-chat/pkg/http/response"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *handler) Chat(w http.ResponseWriter, r *http.Request) {
	chatroomID, _ := uuid.Parse(r.URL.Query().Get("chatroomID"))
	if chatroomID == uuid.Nil {
		err := custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "`chatroomID` is required",
		})

		responser.Error(w, err)
		return
	}

	userID, _ := uuid.Parse(r.URL.Query().Get("userID"))
	if userID == uuid.Nil {
		err := custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "`userID` is required",
		})

		responser.Error(w, err)
		return
	}

	_, err := h.uc.FindOneChatroom(r.Context(), chatroomID)
	if err != nil {
		responser.Error(w, err)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		responser.Error(w, err)
		return
	}

	if _, ok := h.hub.Chatrooms[chatroomID]; !ok {
		h.hub.Chatrooms[chatroomID] = &ws.Chatroom{
			ChatroomID: chatroomID,
			Clients:    make(map[uuid.UUID]*ws.Client),
		}
	}

	client := &ws.Client{
		ID:         userID,
		ChatroomID: chatroomID,
		Conn:       conn,
		MsgChan:    make(chan *ws.Message),
	}

	h.hub.Chatrooms[chatroomID].Clients[client.ID] = client
	h.hub.Register <- client

	go client.Write()
	client.Read(h.hub)
}
