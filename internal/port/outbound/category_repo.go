package outbound

import "kasir-api/internal/model"

type CategoryRepository interface {
	FindAll() ([]model.Category, error)
	FindByID(id int) (*model.Category, error)
	Save(c model.Category) (model.Category, error)
	Update(id int, c model.Category) (*model.Category, error)
	Delete(id int) error
}
