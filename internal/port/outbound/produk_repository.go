package outbound

import (
	"context"
	"kasir-api/internal/domain"
)

type ProdukRepository interface {
	FindAll(ctx context.Context, search string) ([]domain.Produk, error)
	FindByID(ctx context.Context, id int) (domain.Produk, error)
	Save(ctx context.Context, produk domain.Produk) (domain.Produk, error)
	Update(ctx context.Context, produk domain.Produk) (domain.Produk, error)
	Delete(ctx context.Context, id int) error
	UpdateStock(ctx context.Context, id int, qty int) error
}
