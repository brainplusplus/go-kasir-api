package persistence

import (
	"context"
	"kasir-api/internal/port/outbound"

	"gorm.io/gorm"
)

type GormTransactionManager struct {
	db *gorm.DB
}

func NewGormTransactionManager(db *gorm.DB) outbound.TransactionManager {
	return &GormTransactionManager{db: db}
}

type txKey struct{}

func (m *GormTransactionManager) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Inject transaction into context
		ctxWithTx := context.WithValue(ctx, txKey{}, tx)
		return fn(ctxWithTx)
	})
}

// getDB extracts the Gorm transaction from context if available, or returns the standard DB.
// This helper is used by repositories to participate in the transaction.
func getDB(ctx context.Context, standardDB *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return standardDB
}
