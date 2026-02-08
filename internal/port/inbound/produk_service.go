package inbound

import (
	"context"
	"kasir-api/internal/domain"
)

type ProdukService interface {
	GetAll(ctx context.Context, name string) ([]domain.Produk, error)
	GetByID(ctx context.Context, id int) (*domain.Produk, error)
	Create(ctx context.Context, p domain.Produk) (domain.Produk, error)
	Update(ctx context.Context, id int, p domain.Produk) (*domain.Produk, error)
	Delete(ctx context.Context, id int) error
}
