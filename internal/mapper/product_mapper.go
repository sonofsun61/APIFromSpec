package mapper

import (
	"github.com/sonofsun61/APIFromSpec/internal/dto"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
)

func ProductToDTO(product entity.Product) dto.ProductResponse {
	productData := dto.ProductResponse{
		ID: product.ID,
		ProductName: product.ProductName,
		CategoryID: product.CategoryID,
		Price: product.Price,
		AvailableStock: product.AvailableStock,
		SupplierID: product.SupplierID,
	}
	return productData
}

func ProductsToDTO(products []entity.Product) []dto.ProductResponse {
	resp := make([]dto.ProductResponse, len(products))
	for i := range products {
		product := ProductToDTO(products[i])
		resp[i] = product
	}
	return resp
}