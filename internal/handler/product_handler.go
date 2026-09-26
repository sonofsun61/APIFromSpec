package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonofsun61/APIFromSpec/internal/dto"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
	"github.com/sonofsun61/APIFromSpec/internal/mapper"
	"github.com/sonofsun61/APIFromSpec/internal/service"
)

type ProductService interface {
	CreateProduct(ctx context.Context, newProductData entity.NewProductData) (entity.Product, error)
	DecreaseStock(ctx context.Context, productID uuid.UUID, amount int) error
	GetProductByID(ctx context.Context, productID uuid.UUID) (entity.Product, error)
	GetAllProducts(ctx context.Context, limit *int, offset *int) ([]entity.Product, error)
	DeleteProductByID(ctx context.Context, productID uuid.UUID) error
}

type ProductHandler struct {
	service  ProductService
	validate *validator.Validate
}

func NewProductHandler(service ProductService, validate *validator.Validate) *ProductHandler {
	return &ProductHandler{
		service:  service,
		validate: validate,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, "invalid data in JSON", http.StatusBadRequest)
		return
	}
	if req.Price.IsNegative() {
		http.Error(w, "price must not be negative", http.StatusBadRequest)
		return
	}
	newProductData := entity.NewProductData{
		ProductName:    req.ProductName,
		CategoryName:   req.CategoryName,
		Price:          req.Price,
		AvailableStock: req.AvailableStock,
		SupplierID:     req.SupplierID,
	}
	product, err := h.service.CreateProduct(r.Context(), newProductData)
	if err != nil {
		http.Error(w, "failed to create new product", http.StatusInternalServerError)
		return
	}
	resp := dto.ProductResponse{
		ID:             product.ID,
		ProductName:    product.ProductName,
		CategoryID:     product.CategoryID,
		Price:          product.Price,
		AvailableStock: product.AvailableStock,
		SupplierID:     product.SupplierID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *ProductHandler) DecreaseStock(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "failed to parse id", http.StatusBadRequest)
		return
	}
	var req dto.DecreaseAvailableStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if err = h.validate.Struct(req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	err = h.service.DecreaseStock(r.Context(), id, req.Amount)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if errors.Is(err, service.ErrInvalidAmount) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "failed to parse id", http.StatusBadRequest)
		return
	}
	product, err := h.service.GetProductByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
	resp := mapper.ProductToDTO(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var limit, offset *int
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		value, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		limit = &value
	}
	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		value, err := strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		offset = &value
	}
	products, err := h.service.GetAllProducts(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := mapper.ProductsToDTO(products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *ProductHandler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteProductByID(r.Context(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}