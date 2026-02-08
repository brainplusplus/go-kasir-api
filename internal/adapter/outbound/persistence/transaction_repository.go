package persistence

import (
	"context"
	"kasir-api/internal/adapter/outbound/persistence/entity"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/outbound"

	"gorm.io/gorm"
)

type DBTransactionRepository struct {
	db *gorm.DB
}

func NewDBTransactionRepository(db *gorm.DB) outbound.TransactionRepository {
	return &DBTransactionRepository{db: db}
}

func (r *DBTransactionRepository) Create(ctx context.Context, t domain.Transaction) (domain.Transaction, error) {
	var createdTransaction entity.Transaction
	db := getDB(ctx, r.db)

	var details []entity.TransactionDetail
	for _, item := range t.Details {
		details = append(details, entity.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Subtotal:    item.Subtotal,
		})
	}

	createdTransaction = entity.Transaction{
		TotalAmount: t.TotalAmount,
		Details:     details,
	}

	// Just create. Transaction management is verified by service layer.
	// Note: We need to use the db from context (which might be a tx)
	if err := db.Create(&createdTransaction).Error; err != nil {
		return domain.Transaction{}, err
	}

	return createdTransaction.ToDomain(), nil
}

// GetDailyStats removed. Use ReportRepository instead.
