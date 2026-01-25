package memory

import (
	"errors"
	"kasir-api/internal/model"
	"kasir-api/internal/port/outbound"
)

type InMemoryProdukRepository struct {
	produk []model.Produk
}

func NewInMemoryProdukRepository() outbound.ProdukRepository {
	return &InMemoryProdukRepository{
		produk: []model.Produk{
			{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
			{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
			{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
		},
	}
}

func (r *InMemoryProdukRepository) FindAll() ([]model.Produk, error) {
	return r.produk, nil
}

func (r *InMemoryProdukRepository) FindByID(id int) (*model.Produk, error) {
	for _, p := range r.produk {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("produk not found")
}

func (r *InMemoryProdukRepository) Save(p model.Produk) (model.Produk, error) {
	p.ID = len(r.produk) + 1
	r.produk = append(r.produk, p)
	return p, nil
}

func (r *InMemoryProdukRepository) Update(id int, p model.Produk) (*model.Produk, error) {
	for i, existing := range r.produk {
		if existing.ID == id {
			p.ID = id
			r.produk[i] = p
			return &p, nil
		}
	}
	return nil, errors.New("produk not found")
}

func (r *InMemoryProdukRepository) Delete(id int) error {
	for i, p := range r.produk {
		if p.ID == id {
			r.produk = append(r.produk[:i], r.produk[i+1:]...)
			return nil
		}
	}
	return errors.New("produk not found")
}
