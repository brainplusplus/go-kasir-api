package inbound

import "kasir-api/internal/model"

type ProdukService interface {
	GetAll() ([]model.Produk, error)
	GetByID(id int) (*model.Produk, error)
	Create(p model.Produk) (model.Produk, error)
	Update(id int, p model.Produk) (*model.Produk, error)
	Delete(id int) error
}
