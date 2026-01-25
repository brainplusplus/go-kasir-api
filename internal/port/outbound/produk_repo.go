package outbound

import "kasir-api/internal/model"

type ProdukRepository interface {
	FindAll() ([]model.Produk, error)
	FindByID(id int) (*model.Produk, error)
	Save(p model.Produk) (model.Produk, error)
	Update(id int, p model.Produk) (*model.Produk, error)
	Delete(id int) error
}
