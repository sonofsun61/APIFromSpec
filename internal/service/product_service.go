package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, newProductData entity.NewProductData) (entity.Product, error)
	DecreaseStock(ctx context.Context, productID uuid.UUID, amount int) error
	GetProductByID(ctx context.Context, productID uuid.UUID) (entity.Product, error)
	GetAllProducts(ctx context.Context, limit *int, offset *int) ([]entity.Product, error)
	DeleteProductByID(ctx context.Context, productID uuid.UUID) error
}

type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

var ErrInvalidAmount = errors.New("amount must be positive")

func (s *ProductService) CreateProduct(ctx context.Context, newProductData entity.NewProductData) (entity.Product, error) {
	return s.repo.CreateProduct(ctx, newProductData)
}

func (s *ProductService) DecreaseStock(ctx context.Context, productID uuid.UUID, amount int) error {
	if amount > 0 {
		return s.repo.DecreaseStock(ctx, productID, amount)
	}
	return ErrInvalidAmount
}

func (s *ProductService) GetProductByID(ctx context.Context, productID uuid.UUID) (entity.Product, error) {
	return s.repo.GetProductByID(ctx, productID)
}

func (s *ProductService) GetAllProducts(ctx context.Context, limit *int, offset *int) ([]entity.Product, error) {
	return s.repo.GetAllProducts(ctx, limit, offset)
}

func (s *ProductService) DeleteProductByID(ctx context.Context, productID uuid.UUID) error {
	return s.repo.DeleteProductByID(ctx, productID)
}
