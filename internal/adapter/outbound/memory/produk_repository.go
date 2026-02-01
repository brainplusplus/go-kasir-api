package memory

import (
	"errors"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/outbound"
)

type InMemoryProdukRepository struct {
	products []domain.Produk
}

func NewInMemoryProdukRepository() outbound.ProdukRepository {
	return &InMemoryProdukRepository{
		products: []domain.Produk{
			{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10, CategoryID: 1, Category: domain.Category{ID: 1, Name: "Makanan"}},
			{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40, CategoryID: 2, Category: domain.Category{ID: 2, Name: "Minuman"}},
			{ID: 3, Name: "kecap", Price: 12000, Stock: 20, CategoryID: 1, Category: domain.Category{ID: 1, Name: "Makanan"}},
		},
	}
}

func (r *InMemoryProdukRepository) FindAll() ([]domain.Produk, error) {
	return r.products, nil
}

func (r *InMemoryProdukRepository) FindByID(id int) (*domain.Produk, error) {
	for _, p := range r.products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}

func (r *InMemoryProdukRepository) Save(p domain.Produk) (domain.Produk, error) {
	p.ID = len(r.products) + 1
	// In a real memory implementation with join, we would lookup category name here.
	// For simplicity, we assume it's passed or ignored.
	r.products = append(r.products, p)
	return p, nil
}

func (r *InMemoryProdukRepository) Update(id int, p domain.Produk) (*domain.Produk, error) {
	for i, product := range r.products {
		if product.ID == id {
			p.ID = id
			r.products[i] = p
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}

func (r *InMemoryProdukRepository) Delete(id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}
