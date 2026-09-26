package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateProductRequest struct {
	ProductName    string          `json:"product_name" validate:"required,min=1,max=100"`
	CategoryName   string          `json:"category_name" validate:"required,min=1,max=100"`
	Price          decimal.Decimal `json:"price"`
	AvailableStock int             `json:"available_stock" validate:"required,gte=0"`
	SupplierID     uuid.UUID       `json:"supplier_id"`
}

type DecreaseAvailableStockRequest struct {
	Amount int `json:"amount" validate:"gt=0"`
}

type ProductResponse struct {
	ID             uuid.UUID       `json:"id"`
	ProductName    string          `json:"name"`
	CategoryID     uuid.UUID       `json:"category_id"`
	Price          decimal.Decimal `json:"price"`
	AvailableStock int             `json:"available_stock"`
	SupplierID     uuid.UUID       `json:"supplier_id"`
}
