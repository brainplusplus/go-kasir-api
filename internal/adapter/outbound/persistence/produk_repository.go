package persistence

import (
	"errors"
	"kasir-api/internal/adapter/outbound/persistence/entity"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/outbound"

	"gorm.io/gorm"
)

type DBProdukRepository struct {
	db *gorm.DB
}

func NewDBProdukRepository(db *gorm.DB) outbound.ProdukRepository {
	return &DBProdukRepository{db: db}
}

func (r *DBProdukRepository) FindAll() ([]domain.Produk, error) {
	var products []entity.Produk
	if err := r.db.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	domainProducts := make([]domain.Produk, 0)
	for _, p := range products {
		domainProducts = append(domainProducts, p.ToDomain())
	}
	return domainProducts, nil
}

func (r *DBProdukRepository) FindByID(id int) (*domain.Produk, error) {
	var product entity.Produk
	if err := r.db.Preload("Category").First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	d := product.ToDomain()
	return &d, nil
}

func (r *DBProdukRepository) Save(p domain.Produk) (domain.Produk, error) {
	product := entity.FromDomainProduk(p)
	if err := r.db.Create(&product).Error; err != nil {
		return domain.Produk{}, err
	}

	// Reload with Category
	if err := r.db.Preload("Category").First(&product, product.ID).Error; err != nil {
		// Log error but maybe don't fail? Or fail because consistency is key?
		// Failing here is safer to indicate something went wrong with fetching back.
		return domain.Produk{}, err
	}

	return product.ToDomain(), nil
}

func (r *DBProdukRepository) Update(id int, p domain.Produk) (*domain.Produk, error) {
	var product entity.Produk
	if err := r.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Stock = p.Stock
	product.CategoryID = p.CategoryID // Ensure CategoryID is updated if provided

	if err := r.db.Save(&product).Error; err != nil {
		return nil, err
	}

	// Reload with Category
	if err := r.db.Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}

	d := product.ToDomain()
	return &d, nil
}

func (r *DBProdukRepository) Delete(id int) error {
	result := r.db.Delete(&entity.Produk{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}
