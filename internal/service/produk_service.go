package service

import (
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

func (s *productService) GetAll() ([]domain.Produk, error) {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id int) (*domain.Produk, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Create(p domain.Produk) (domain.Produk, error) {
	return s.repo.Save(p)
}

func (s *productService) Update(id int, p domain.Produk) (*domain.Produk, error) {
	return s.repo.Update(id, p)
}

func (s *productService) Delete(id int) error {
	return s.repo.Delete(id)
}
