package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonofsun61/APIFromSpec/internal/dto"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
)

type ClientService interface {
	CreateClient(ctx context.Context, newClientData entity.NewClientData, country string, city string, street string) (uuid.UUID, error)
	DeleteClientByID(ctx context.Context, clientID uuid.UUID) error
	GetClientByNameAndSurname(ctx context.Context, name string, surname string) ([]entity.ClientWithAddress, error)
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

func (h *ClientHandler) DeleteClientByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteClientByID(r.Context(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ClientHandler) GetClientByNameAndSurname(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	clients, err := h.service.GetClientByNameAndSurname(r.Context(), name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := make([]dto.ClientResponse, len(clients))
	for i := 0; i < len(clients); i++ {
		address := dto.AddressResponse{
			Country: clients[i].Country,
			City: clients[i].City,
			Street: clients[i].Street,
		}
		client := dto.ClientResponse{
			ID: clients[i].ID,
			ClientName: clients[i].ClientName,
			ClientSurname: clients[i].ClientSurname,
			Birthday: clients[i].Birthday,
			Gender: clients[i].Gender,
			Address: address,
		}
		resp[i] = client
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}