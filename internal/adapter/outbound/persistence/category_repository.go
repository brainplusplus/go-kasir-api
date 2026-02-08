package persistence

import (
	"context"
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

func (r *DBCategoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	var categories []entity.Category
	db := getDB(ctx, r.db)
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	domainCategories := make([]domain.Category, 0)
	for _, c := range categories {
		domainCategories = append(domainCategories, c.ToDomain())
	}
	return domainCategories, nil
}

func (r *DBCategoryRepository) FindByID(ctx context.Context, id int) (*domain.Category, error) {
	var category entity.Category
	db := getDB(ctx, r.db)
	if err := db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	d := category.ToDomain()
	return &d, nil
}

func (r *DBCategoryRepository) Save(ctx context.Context, c domain.Category) (domain.Category, error) {
	category := entity.FromDomainCategory(c)
	db := getDB(ctx, r.db)
	if err := db.Create(&category).Error; err != nil {
		return domain.Category{}, err
	}
	return category.ToDomain(), nil
}

func (r *DBCategoryRepository) Update(ctx context.Context, id int, c domain.Category) (*domain.Category, error) {
	var category entity.Category
	db := getDB(ctx, r.db)
	if err := db.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	// Update fields
	category.Name = c.Name
	category.Description = c.Description

	if err := db.Save(&category).Error; err != nil {
		return nil, err
	}

	d := category.ToDomain()
	return &d, nil
}

func (r *DBCategoryRepository) Delete(ctx context.Context, id int) error {
	db := getDB(ctx, r.db)
	result := db.Delete(&entity.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}
