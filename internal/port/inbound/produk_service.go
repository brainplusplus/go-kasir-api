package inbound

import "kasir-api/internal/domain"

type ProdukService interface {
	GetAll() ([]domain.Produk, error)
	GetByID(id int) (*domain.Produk, error)
	Create(p domain.Produk) (domain.Produk, error)
	Update(id int, p domain.Produk) (*domain.Produk, error)
	Delete(id int) error
}
