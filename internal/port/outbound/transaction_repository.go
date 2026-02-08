package outbound

import (
	"context"
	"kasir-api/internal/domain"
)

type TransactionRepository interface {
	Create(ctx context.Context, t domain.Transaction) (domain.Transaction, error)
}
