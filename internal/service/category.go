package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/port/outbound"
)

type categoryService struct {
	repo outbound.CategoryRepository
}

func NewCategoryService(repo outbound.CategoryRepository) inbound.CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) GetByID(id int) (*model.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) Create(c model.Category) (model.Category, error) {
	return s.repo.Save(c)
}

func (s *categoryService) Update(id int, c model.Category) (*model.Category, error) {
	return s.repo.Update(id, c)
}

func (s *categoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
