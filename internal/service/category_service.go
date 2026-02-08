package service

import (
	"context"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/port/outbound"
)

type categoryService struct {
	repo outbound.CategoryRepository
}

func NewCategoryService(repo outbound.CategoryRepository) inbound.CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	return s.repo.FindAll(ctx)
}

func (s *categoryService) GetByID(ctx context.Context, id int) (*domain.Category, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *categoryService) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
	return s.repo.Save(ctx, c)
}

func (s *categoryService) Update(ctx context.Context, id int, c domain.Category) (*domain.Category, error) {
	return s.repo.Update(ctx, id, c)
}

func (s *categoryService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
