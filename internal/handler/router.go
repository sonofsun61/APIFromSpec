package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetUpRouter(clientHandler *ClientHandler, supplierHandler *SupplierHandler, productHandler *ProductHandler, imageHandler *ImageHandler) http.Handler {
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
		r.Route("/products", func(r chi.Router) {
			r.Post("/", productHandler.CreateProduct)
			r.Patch("/{id}/stock", productHandler.DecreaseStock)
			r.Get("/{id}", productHandler.GetProductByID)
			r.Get("/", productHandler.GetAllProducts)
			r.Delete("/{id}", productHandler.DeleteProductByID)
			r.Post("/{id}/image", imageHandler.AddImage)
			r.Get("/{id}/image", imageHandler.GetImageByProductID)
		})
		r.Route("/images", func(r chi.Router) {
			r.Patch("/{id}", imageHandler.ChangeImage)
			r.Delete("/{id}", imageHandler.DeleteImage)
			r.Get("/{id}", imageHandler.GetImageByImageID)
		})
	})
	return r
}