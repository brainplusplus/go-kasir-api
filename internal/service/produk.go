package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/port/outbound"
)

type productService struct {
	repo outbound.ProdukRepository
}

func NewProductService(repo outbound.ProdukRepository) inbound.ProdukService {
	return &productService{repo: repo}
}

func (s *productService) GetAll() ([]model.Produk, error) {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id int) (*model.Produk, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Create(p model.Produk) (model.Produk, error) {
	return s.repo.Save(p)
}

func (s *productService) Update(id int, p model.Produk) (*model.Produk, error) {
	return s.repo.Update(id, p)
}

func (s *productService) Delete(id int) error {
	return s.repo.Delete(id)
}
