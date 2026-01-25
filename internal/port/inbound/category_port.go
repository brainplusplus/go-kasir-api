package inbound

import "kasir-api/internal/model"

type CategoryService interface {
	GetAll() ([]model.Category, error)
	GetByID(id int) (*model.Category, error)
	Create(c model.Category) (model.Category, error)
	Update(id int, c model.Category) (*model.Category, error)
	Delete(id int) error
}
