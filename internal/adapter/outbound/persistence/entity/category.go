package entity

import (
	"kasir-api/internal/domain"
	"time"
)

type Category struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (c *Category) ToDomain() domain.Category {
	return domain.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}

func FromDomainCategory(c domain.Category) Category {
	return Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
