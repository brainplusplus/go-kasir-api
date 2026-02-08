package inbound

import (
	"context"
	"kasir-api/internal/domain"
)

type CategoryService interface {
	GetAll(ctx context.Context) ([]domain.Category, error)
	GetByID(ctx context.Context, id int) (*domain.Category, error)
	Create(ctx context.Context, c domain.Category) (domain.Category, error)
	Update(ctx context.Context, id int, c domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, id int) error
}
