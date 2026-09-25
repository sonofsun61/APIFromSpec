package dto

import "github.com/google/uuid"

type SupplierRequest struct {
	SupplierName string `json:"name" validate:"required,min=1,max=100"`
	PhoneNumber  string `json:"phone_number" validate:"required,e164"`
	Country      string `json:"country" validate:"required,min=1,max=25"`
	City         string `json:"city" validate:"required,min=1,max=50"`
	Street       string `json:"street" validate:"required,min=1,max=50"`
}

type SupplierResponse struct {
	ID           uuid.UUID       `json:"id"`
	SupplierName string          `json:"name"`
	PhoneNumber  string          `json:"phone_number"`
	Address      AddressResponse `json:"address"`
}
