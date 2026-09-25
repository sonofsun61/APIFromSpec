package dto

type UpdateAddressRequest struct {
	Country string `json:"country" validate:"required,min=1,max=25"`
	City    string `json:"city" validate:"required,min=1,max=50"`
	Street  string `json:"street" validate:"required,min=1,max=50"`
}

type AddressResponse struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}
