package memory

import (
	"context"
	"errors"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/outbound"
	"strings"
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

func (r *InMemoryProdukRepository) FindAll(ctx context.Context, search string) ([]domain.Produk, error) {
	if search == "" {
		return r.products, nil
	}
	var filtered []domain.Produk
	for _, p := range r.products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(search)) {
			filtered = append(filtered, p)
		}
	}
	return filtered, nil
}

func (r *InMemoryProdukRepository) FindByID(ctx context.Context, id int) (domain.Produk, error) {
	for _, p := range r.products {
		if p.ID == id {
			return p, nil
		}
	}
	return domain.Produk{}, errors.New("product not found")
}

func (r *InMemoryProdukRepository) Save(ctx context.Context, p domain.Produk) (domain.Produk, error) {
	p.ID = len(r.products) + 1
	r.products = append(r.products, p)
	return p, nil
}

func (r *InMemoryProdukRepository) Update(ctx context.Context, p domain.Produk) (domain.Produk, error) {
	for i, product := range r.products {
		if product.ID == p.ID {
			// Update fields
			// Ensure we keep ID and maybe verify existence
			r.products[i] = p
			return p, nil
		}
	}
	return domain.Produk{}, errors.New("product not found")
}

func (r *InMemoryProdukRepository) Delete(ctx context.Context, id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

func (r *InMemoryProdukRepository) UpdateStock(ctx context.Context, id int, qty int) error {
	for i, p := range r.products {
		if p.ID == id {
			if p.Stock < qty {
				return errors.New("insufficient stock")
			}
			r.products[i].Stock -= qty
			return nil
		}
	}
	return errors.New("product not found")
}
