package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, clientData entity.Client) (uuid.UUID, error)
	DeleteClientByID(ctx context.Context, clientID uuid.UUID) error
	GetClientByNameAndSurname(ctx context.Context, name string, surname string) ([]entity.Client, error)
	GetAllClients(ctx context.Context, limit *int, offset *int) ([]entity.Client, error)
	UpdateClientAddress(ctx context.Context, clientID uuid.UUID, country string, city string, street string) error
}

type ClientService struct {
	repo ClientRepository
}

func NewClientService(repo ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) CreateClient(ctx context.Context, clientData entity.Client) (uuid.UUID, error) {
	return s.repo.CreateClient(ctx, clientData)
}

func (s *ClientService) DeleteClientByID(ctx context.Context, clientID uuid.UUID) error {
	return s.repo.DeleteClientByID(ctx, clientID)
}

func (s *ClientService) GetClientByNameAndSurname(ctx context.Context, name string, surname string) ([]entity.Client, error) {
	return s.repo.GetClientByNameAndSurname(ctx, name, surname)
}

func (s *ClientService) GetAllClients(ctx context.Context, limit *int, offset *int) ([]entity.Client, error) {
	return s.repo.GetAllClients(ctx, limit, offset)
}

func (s *ClientService) UpdateClientAddress(ctx context.Context, clientID uuid.UUID, country string, city string, street string) error {
	return s.repo.UpdateClientAddress(ctx, clientID, country, city, street)
}