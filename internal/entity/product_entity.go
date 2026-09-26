package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID             uuid.UUID       `db:"id"`
	ProductName    string          `db:"name"`
	CategoryID     uuid.UUID       `db:"category_id"`
	Price          decimal.Decimal `db:"price"`
	AvailableStock int             `db:"available_stock"`
	LastUpdateDate time.Time       `db:"last_update_date"`
	SupplierID     uuid.UUID       `db:"supplier_id"`
	ImageID        *uuid.UUID      `db:"image_id"`
}

type NewProductData struct {
	ProductName    string
	CategoryName   string
	Price          decimal.Decimal
	AvailableStock int
	SupplierID     uuid.UUID
}
