package outbound

import (
	"context"
	"kasir-api/internal/domain"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]domain.Category, error)
	FindByID(ctx context.Context, id int) (*domain.Category, error)
	Save(ctx context.Context, c domain.Category) (domain.Category, error)
	Update(ctx context.Context, id int, c domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, id int) error
}
