package transporter

import (
	"encoding/json"
	"net/http"

	"github.com/ffajarpratama/go-chat/internal/handler"
	"github.com/ffajarpratama/go-chat/internal/usecase"
	"github.com/ffajarpratama/go-chat/pkg/constant"
	base_response "github.com/ffajarpratama/go-chat/pkg/http/response"
	custom_validator "github.com/ffajarpratama/go-chat/pkg/validator"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	go_validator "github.com/go-playground/validator/v10"
)

func NewV1HttpTransporter(uc *usecase.Usecase) http.Handler {
	r := chi.NewRouter()
	v := custom_validator.New(go_validator.New())

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(base_response.JsonResponse{
			Error: &base_response.ErrorResponse{
				Code:    constant.DefaultNotFoundError,
				Status:  http.StatusNotFound,
				Message: "please check url",
			},
		})
	})

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(base_response.JsonResponse{
			Success: true,
			Data: map[string]interface{}{
				"app-name": "go-chat",
			},
		})
	})

	r.Mount("/api", handler.New(uc, v))

	return r
}
