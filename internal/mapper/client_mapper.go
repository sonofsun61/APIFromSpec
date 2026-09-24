package mapper

import (
	"github.com/sonofsun61/APIFromSpec/internal/dto"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
)

func ClientWithAddressToDTO(client entity.ClientWithAddress) dto.ClientResponse {
	address := dto.AddressResponse{
			Country: client.Country,
			City: client.City,
			Street: client.Street,
		}
	clientData := dto.ClientResponse{
		ID: client.ID,
		ClientName: client.ClientName,
		ClientSurname: client.ClientSurname,
		Birthday: client.Birthday,
		Gender: client.Gender,
		Address: address,
	}
	return clientData
}

func ClientsWithAddressesToDTO(clients []entity.ClientWithAddress) []dto.ClientResponse {
	resp := make([]dto.ClientResponse, len(clients))
	for i := range clients {
		client := ClientWithAddressToDTO(clients[i])
		resp[i] = client
	}
	return resp
}