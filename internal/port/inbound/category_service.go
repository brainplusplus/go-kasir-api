package inbound

import "kasir-api/internal/domain"

type CategoryService interface {
	GetAll() ([]domain.Category, error)
	GetByID(id int) (*domain.Category, error)
	Create(c domain.Category) (domain.Category, error)
	Update(id int, c domain.Category) (*domain.Category, error)
	Delete(id int) error
}
