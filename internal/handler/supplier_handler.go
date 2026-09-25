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
)

type SupplierService interface {
	AddSupplier(ctx context.Context, newSupplierData entity.NewSupplierData, country string, city string, street string) (uuid.UUID, error)
	UpdateSupplierAddress(ctx context.Context, supplierID uuid.UUID, country string, city string, street string) error
	DeleteSupplier(ctx context.Context, supplierID uuid.UUID) error
	GetSuppliers(ctx context.Context, limit *int, offset *int) ([]entity.SupplierWithAddress, error)
	GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (entity.SupplierWithAddress, error)
}

type SupplierHandler struct {
	service  SupplierService
	validate *validator.Validate
}

func NewSupplierHandler(service SupplierService, validate *validator.Validate) *SupplierHandler {
	return &SupplierHandler{
		service:  service,
		validate: validate,
	}
}

func (h *SupplierHandler) AddSupplier(w http.ResponseWriter, r *http.Request) {
	var req dto.SupplierRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newSupplierData := entity.NewSupplierData{
		SupplierName: req.SupplierName,
		PhoneNumber:  req.PhoneNumber,
	}
	newSupplierID, err := h.service.AddSupplier(r.Context(), newSupplierData, req.Country, req.City, req.Street)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.SupplierResponse{
		ID: newSupplierID,
		SupplierName: req.SupplierName,
		PhoneNumber: req.PhoneNumber,
		Address: dto.AddressResponse{
			Country: req.Country,
			City: req.City,
			Street: req.Street,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *SupplierHandler) UpdateSupplierAddress(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req dto.UpdateAddressRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateSupplierAddress(r.Context(), id, req.Country, req.City, req.Street); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *SupplierHandler) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteSupplier(r.Context(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *SupplierHandler) GetSuppliers(w http.ResponseWriter, r *http.Request) {
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
	suppliers, err := h.service.GetSuppliers(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := mapper.SupplierWithAddressesToDTO(suppliers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *SupplierHandler) GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	supplierData, err := h.service.GetSupplierByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := mapper.SupplierWithAddressToDTO(supplierData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
