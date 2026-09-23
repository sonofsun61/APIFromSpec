package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateClientRequest struct {
	ClientName    string    `json:"client_name" validate:"required,min=1,max=100"`
	ClientSurname string    `json:"client_surname" validate:"required,min=1,max=100"`
	Birthday      time.Time `json:"birthday" validate:"required,lt"`
	Gender        string    `json:"gender" validate:"oneof=male female"`
	Country       string    `json:"country" validate:"required,min=1,max=25"`
	City          string    `json:"city" validate:"required,min=1,max=50"`
	Street        string    `json:"street" validate:"required,min=1,max=50"`
}

type ClientResponse struct {
	ID            uuid.UUID       `json:"id"`
	ClientName    string          `json:"client_name"`
	ClientSurname string          `json:"client_surname"`
	Birthday      time.Time       `json:"birthday"`
	Gender        string          `json:"gender"`
	Address       AddressResponse `json:"address"`
}
