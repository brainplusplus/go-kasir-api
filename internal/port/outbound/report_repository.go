package outbound

import (
	"context"
	"kasir-api/internal/domain"
	"time"
)

type ReportRepository interface {
	GetStatsByRange(ctx context.Context, startDate, endDate time.Time) (domain.Report, error)
}
