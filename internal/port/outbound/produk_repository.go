package outbound

import "kasir-api/internal/domain"

type ProdukRepository interface {
	FindAll() ([]domain.Produk, error)
	FindByID(id int) (*domain.Produk, error)
	Save(p domain.Produk) (domain.Produk, error)
	Update(id int, p domain.Produk) (*domain.Produk, error)
	Delete(id int) error
}
