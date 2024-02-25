package handler

import (
	"net/http"

	"github.com/ffajarpratama/go-chat/internal/usecase"
	"github.com/ffajarpratama/go-chat/internal/ws"
	custom_validator "github.com/ffajarpratama/go-chat/pkg/validator"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	uc  *usecase.Usecase
	v   *custom_validator.Validator
	hub *ws.Hub
}

func New(uc *usecase.Usecase, v *custom_validator.Validator) http.Handler {
	r := chi.NewRouter()
	hub := ws.NewHub()
	go hub.Run()

	h := &handler{
		uc:  uc,
		v:   v,
		hub: hub,
	}

	r.Route("/user", func(user chi.Router) {
		user.Get("/", h.FindAndCountUser)
	})

	r.Route("/chatroom", func(chatroom chi.Router) {
		chatroom.Post("/", h.CreateChatroom)
		chatroom.Get("/", h.FindAndCountChatroom)
	})

	r.Route("/chat", func(r chi.Router) {
		r.Get("/", h.Chat)
	})

	return r
}
