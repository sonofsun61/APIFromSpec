package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetUpRouter(clientHandler *ClientHandler, supplierHandler *SupplierHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router)  {
		r.Route("/clients", func(r chi.Router) {
			r.Post("/", clientHandler.CreateClient)
			r.Delete("/{id}", clientHandler.DeleteClientByID)
			r.Get("/", clientHandler.GetClients)
			r.Patch("/{id}", clientHandler.UpdateClientAddress)
		})
		r.Route("/suppliers", func(r chi.Router) {
			r.Post("/", supplierHandler.AddSupplier)
			r.Patch("/{id}", supplierHandler.UpdateSupplierAddress)
			r.Delete("/{id}", supplierHandler.DeleteSupplier)
			r.Get("/", supplierHandler.GetSuppliers)
			r.Get("/{id}", supplierHandler.GetSupplierByID)
		})
	})
	return r
}