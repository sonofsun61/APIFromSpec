package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetUpRouter(clientHandler *ClientHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router)  {
		r.Route("/clients", func(r chi.Router) {
			r.Post("/", clientHandler.CreateClient)
		})
	})
	return r
}