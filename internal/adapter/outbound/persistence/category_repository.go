package persistence

import (
	"errors"
	"kasir-api/internal/adapter/outbound/persistence/entity"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/outbound"

	"gorm.io/gorm"
)

type DBCategoryRepository struct {
	db *gorm.DB
}

func NewDBCategoryRepository(db *gorm.DB) outbound.CategoryRepository {
	return &DBCategoryRepository{db: db}
}

func (r *DBCategoryRepository) FindAll() ([]domain.Category, error) {
	var categories []entity.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	domainCategories := make([]domain.Category, 0)
	for _, c := range categories {
		domainCategories = append(domainCategories, c.ToDomain())
	}
	return domainCategories, nil
}

func (r *DBCategoryRepository) FindByID(id int) (*domain.Category, error) {
	var category entity.Category
	if err := r.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	d := category.ToDomain()
	return &d, nil
}

func (r *DBCategoryRepository) Save(c domain.Category) (domain.Category, error) {
	category := entity.FromDomainCategory(c)
	if err := r.db.Create(&category).Error; err != nil {
		return domain.Category{}, err
	}
	return category.ToDomain(), nil
}

func (r *DBCategoryRepository) Update(id int, c domain.Category) (*domain.Category, error) {
	var category entity.Category
	if err := r.db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	// Update fields
	category.Name = c.Name
	category.Description = c.Description

	if err := r.db.Save(&category).Error; err != nil {
		return nil, err
	}

	d := category.ToDomain()
	return &d, nil
}

func (r *DBCategoryRepository) Delete(id int) error {
	result := r.db.Delete(&entity.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}
