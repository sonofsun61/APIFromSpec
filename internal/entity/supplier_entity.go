package entity

import "github.com/google/uuid"

type Supplier struct {
	ID           uuid.UUID `db:"id"`
	SupplierName string    `db:"name"`
	PhoneNumber  string    `db:"phone_number"`
	AddressID    uuid.UUID `db:"address_id"`
}

type NewSupplierData struct {
	SupplierName string
	PhoneNumber  string
}

type SupplierWithAddress struct {
	ID           uuid.UUID `db:"id"`
	SupplierName string    `db:"name"`
	PhoneNumber  string    `db:"phone_number"`
	Country      string    `db:"country"`
	City         string    `db:"city"`
	Street       string    `db:"street"`
}
