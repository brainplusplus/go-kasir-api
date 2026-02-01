package dto

import "kasir-api/internal/domain"

type CreateProdukRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

type UpdateProdukRequest struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

type ProdukResponse struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Price      int              `json:"price"`
	Stock      int              `json:"stock"`
	CategoryID int              `json:"category_id"`
	Category   CategoryResponse `json:"category"`
}

func (r *CreateProdukRequest) ToDomain() domain.Produk {
	return domain.Produk{
		Name:       r.Name,
		Price:      r.Price,
		Stock:      r.Stock,
		CategoryID: r.CategoryID,
	}
}

func (r *UpdateProdukRequest) ToDomain() domain.Produk {
	return domain.Produk{
		Name:       r.Name,
		Price:      r.Price,
		Stock:      r.Stock,
		CategoryID: r.CategoryID,
	}
}

func FromDomainProduk(p domain.Produk) ProdukResponse {
	return ProdukResponse{
		ID:         p.ID,
		Name:       p.Name,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryID,
		Category: CategoryResponse{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		},
	}
}
