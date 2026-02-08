package persistence

import (
	"context"
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

func (r *DBProdukRepository) FindAll(ctx context.Context, search string) ([]domain.Produk, error) {
	var produkEntities []entity.Produk
	db := getDB(ctx, r.db)

	query := db.Preload("Category")
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if err := query.Find(&produkEntities).Error; err != nil {
		return nil, err
	}

	var produks []domain.Produk
	for _, p := range produkEntities {
		produks = append(produks, p.ToDomain())
	}
	return produks, nil
}

func (r *DBProdukRepository) FindByID(ctx context.Context, id int) (domain.Produk, error) {
	var produkEntity entity.Produk
	db := getDB(ctx, r.db)
	if err := db.Preload("Category").First(&produkEntity, id).Error; err != nil {
		return domain.Produk{}, err
	}
	return produkEntity.ToDomain(), nil
}

func (r *DBProdukRepository) Save(ctx context.Context, produk domain.Produk) (domain.Produk, error) {
	produkEntity := entity.FromDomainProduk(produk)
	db := getDB(ctx, r.db)
	if err := db.Create(&produkEntity).Error; err != nil {
		return domain.Produk{}, err
	}
	return produkEntity.ToDomain(), nil
}

func (r *DBProdukRepository) Update(ctx context.Context, produk domain.Produk) (domain.Produk, error) {
	produkEntity := entity.FromDomainProduk(produk)
	db := getDB(ctx, r.db)
	// Make sure to use ID for update
	if err := db.Model(&entity.Produk{}).Where("id = ?", produk.ID).Updates(produkEntity).Error; err != nil {
		return domain.Produk{}, err
	}
	// Fetch updated to return complete object including associations if needed
	return r.FindByID(ctx, produk.ID)
}

func (r *DBProdukRepository) Delete(ctx context.Context, id int) error {
	db := getDB(ctx, r.db)
	if err := db.Delete(&entity.Produk{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *DBProdukRepository) UpdateStock(ctx context.Context, id int, qty int) error {
	db := getDB(ctx, r.db)
	// Decrement stock atomically
	// Note: qty is the amount to SUBTRACT
	result := db.Model(&entity.Produk{}).
		Where("id = ? AND stock >= ?", id, qty).
		UpdateColumn("stock", gorm.Expr("stock - ?", qty))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insufficient stock or product not found")
	}
	return nil
}
