package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonofsun61/APIFromSpec/internal/entity"
)

type ImageRepository interface {
	AddImage(ctx context.Context, imageBytes []byte, productID uuid.UUID) error
	ChangeImage(ctx context.Context, imageID uuid.UUID, imageBytes []byte) error
	DeleteImage(ctx context.Context, imageID uuid.UUID) error
	GetImageByProductID(ctx context.Context, productID uuid.UUID) (entity.Image, error)
	GetImageByImageID(ctx context.Context, imageID uuid.UUID) (entity.Image, error)
}

type ImageService struct {
	repo ImageRepository
}

func NewImageService(repo ImageRepository) *ImageService {
	return &ImageService{
		repo: repo,
	}
}
func (s *ImageService) AddImage(ctx context.Context, bytes []byte, productID uuid.UUID) error {
	return s.repo.AddImage(ctx, bytes, productID)
}

func (s *ImageService) ChangeImage(ctx context.Context, imageID uuid.UUID, bytes []byte) error {
	return s.repo.ChangeImage(ctx, imageID, bytes)
}

func (s *ImageService) DeleteImage(ctx context.Context, imageID uuid.UUID) error {
	return s.repo.DeleteImage(ctx, imageID)
}

func (s *ImageService) GetImageByProductID(ctx context.Context, productID uuid.UUID) (entity.Image, error) {
	return s.repo.GetImageByProductID(ctx, productID)
}

func (s *ImageService) GetImageByImageID(ctx context.Context, imageID uuid.UUID) (entity.Image, error) {
	return s.repo.GetImageByImageID(ctx, imageID)
}
