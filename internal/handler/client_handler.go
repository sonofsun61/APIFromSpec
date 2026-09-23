package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sonofsun61/APIFromSpec/internal/dto"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
)

type ClientService interface {
	CreateClient(ctx context.Context, newClientData entity.NewClientData, country string, city string, street string) (uuid.UUID, error)
	DeleteClientByID(ctx context.Context, clientID uuid.UUID) error
	GetClientByNameAndSurname(ctx context.Context, name string, surname string) ([]entity.Client, error)
	GetAllClients(ctx context.Context, limit *int, offset *int) ([]entity.Client, error)
	UpdateClientAddress(ctx context.Context, clientID uuid.UUID, country string, city string, street string) error
}

type ClientHandler struct {
	service  ClientService
	validate *validator.Validate
}

func NewClientHandler(service ClientService, validate *validator.Validate) *ClientHandler {
	return &ClientHandler{
		service:  service,
		validate: validate,
	}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newClientData := entity.NewClientData{
		ClientName: req.ClientName,
		ClientSurname: req.ClientSurname,
		Birthday: req.Birthday,
		Gender: req.Gender,
	}
	newID, err := h.service.CreateClient(r.Context(), newClientData, req.Country, req.City, req.Street)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := dto.ClientResponse{
		ID: newID,
		ClientName: req.ClientName,
		ClientSurname: req.ClientSurname,
		Birthday: req.Birthday,
		Gender: req.Gender,
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