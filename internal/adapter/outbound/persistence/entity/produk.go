package entity

import (
	"kasir-api/internal/domain"
	"time"
)

type Produk struct {
	ID         int       `gorm:"primaryKey"`
	Name       string    `gorm:"column:name"`
	Price      int       `gorm:"column:price"`
	Stock      int       `gorm:"column:stock"`
	CategoryID int       `gorm:"column:category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (p *Produk) ToDomain() domain.Produk {
	return domain.Produk{
		ID:         p.ID,
		Name:       p.Name,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryID,
		Category: domain.Category{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		},
	}
}

func FromDomainProduk(p domain.Produk) Produk {
	return Produk{
		ID:         p.ID,
		Name:       p.Name,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryID,
	}
}
