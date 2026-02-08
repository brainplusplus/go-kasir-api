package inbound

import (
	"context"
	"kasir-api/internal/adapter/inbound/http/dto"
	"kasir-api/internal/domain"
	"time"
)

type TransactionService interface {
	Checkout(ctx context.Context, request dto.CheckoutRequest) (domain.Transaction, error)
	GetReport(ctx context.Context, startDate, endDate time.Time) (domain.Report, error)
	GetReportToday(ctx context.Context) (domain.Report, error)
}
