package memory

import (
	"errors"
	"kasir-api/internal/model"
	"kasir-api/internal/port/outbound"
)

type InMemoryCategoryRepository struct {
	category []model.Category
}

func NewInMemoryCategoryRepository() outbound.CategoryRepository {
	return &InMemoryCategoryRepository{
		category: []model.Category{
			{ID: 1, Name: "Makanan", Description: ""},
			{ID: 2, Name: "Minuman", Description: ""},
			{ID: 3, Name: "Alat Tulis", Description: ""},
		},
	}
}

func (r *InMemoryCategoryRepository) FindAll() ([]model.Category, error) {
	return r.category, nil
}

func (r *InMemoryCategoryRepository) FindByID(id int) (*model.Category, error) {
	for _, c := range r.category {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("category not found")
}

func (r *InMemoryCategoryRepository) Save(c model.Category) (model.Category, error) {
	c.ID = len(r.category) + 1
	r.category = append(r.category, c)
	return c, nil
}

func (r *InMemoryCategoryRepository) Update(id int, c model.Category) (*model.Category, error) {
	for i, existing := range r.category {
		if existing.ID == id {
			c.ID = id
			r.category[i] = c
			return &c, nil
		}
	}
	return nil, errors.New("category not found")
}

func (r *InMemoryCategoryRepository) Delete(id int) error {
	for i, c := range r.category {
		if c.ID == id {
			r.category = append(r.category[:i], r.category[i+1:]...)
			return nil
		}
	}
	return errors.New("category not found")
}
