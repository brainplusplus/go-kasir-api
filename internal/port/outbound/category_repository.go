package outbound

import "kasir-api/internal/domain"

type CategoryRepository interface {
	FindAll() ([]domain.Category, error)
	FindByID(id int) (*domain.Category, error)
	Save(c domain.Category) (domain.Category, error)
	Update(id int, c domain.Category) (*domain.Category, error)
	Delete(id int) error
}
