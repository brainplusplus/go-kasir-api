package service

import (
	"context"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/port/outbound"
)

type productService struct {
	repo outbound.ProdukRepository
}

func NewProdukService(repo outbound.ProdukRepository) inbound.ProdukService {
	return &productService{repo: repo}
}

func (s *productService) GetAll(ctx context.Context, name string) ([]domain.Produk, error) {
	return s.repo.FindAll(ctx, name)
}

func (s *productService) GetByID(ctx context.Context, id int) (*domain.Produk, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *productService) Create(ctx context.Context, p domain.Produk) (domain.Produk, error) {
	return s.repo.Save(ctx, p)
}

func (s *productService) Update(ctx context.Context, id int, p domain.Produk) (*domain.Produk, error) {
	// Ensure ID is set
	p.ID = id
	updated, err := s.repo.Update(ctx, p)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (s *productService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
