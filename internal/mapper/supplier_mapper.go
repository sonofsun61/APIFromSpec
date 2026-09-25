package mapper

import (
	"github.com/sonofsun61/APIFromSpec/internal/dto"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
)

func SupplierWithAddressToDTO(supplier entity.SupplierWithAddress) dto.SupplierResponse {
	address := dto.AddressResponse{
			Country: supplier.Country,
			City: supplier.City,
			Street: supplier.Street,
		}
	supplierData := dto.SupplierResponse{
		ID: supplier.ID,
		SupplierName: supplier.SupplierName,
		PhoneNumber: supplier.PhoneNumber,
		Address: address,
	}
	return supplierData
}

func SupplierWithAddressesToDTO(suppliers []entity.SupplierWithAddress) []dto.SupplierResponse {
	resp := make([]dto.SupplierResponse, len(suppliers))
	for i := range suppliers {
		supplier := SupplierWithAddressToDTO(suppliers[i])
		resp[i] = supplier
	}
	return resp
}