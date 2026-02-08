package persistence

import (
	"context"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/outbound"
	"time"

	"gorm.io/gorm"
)

type DBReportRepository struct {
	db *gorm.DB
}

func NewDBReportRepository(db *gorm.DB) outbound.ReportRepository {
	return &DBReportRepository{db: db}
}

func (r *DBReportRepository) GetStatsByRange(ctx context.Context, startDate, endDate time.Time) (domain.Report, error) {
	var report domain.Report
	db := getDB(ctx, r.db)

	// Single Query Optimization
	// We use subqueries/CTEs to fetch all metrics in one go.
	// Note: GORM's Raw can map to struct fields with `scan`.
	type Result struct {
		TotalRevenue      int
		TotalTransactions int
		TopProductName    string
		TopProductQty     int
	}
	var res Result

	query := `
		SELECT 
			COALESCE(SUM(t.total_amount), 0) as total_revenue,
			COUNT(t.id) as total_transactions,
			(
				SELECT td.product_name 
				FROM transaction_details td
				JOIN transactions t2 ON t2.id = td.transaction_id
				WHERE t2.created_at >= ? AND t2.created_at < ?
				GROUP BY td.product_name 
				ORDER BY SUM(td.qty) DESC 
				LIMIT 1
			) as top_product_name,
			(
				SELECT SUM(td.qty) 
				FROM transaction_details td
				JOIN transactions t2 ON t2.id = td.transaction_id
				WHERE t2.created_at >= ? AND t2.created_at < ?
				GROUP BY td.product_name 
				ORDER BY SUM(td.qty) DESC 
				LIMIT 1
			) as top_product_qty
		FROM transactions t
		WHERE t.created_at >= ? AND t.created_at < ?
	`

	// Pass arguments 3 times (main query + 2 subqueries)
	if err := db.Raw(query, startDate, endDate, startDate, endDate, startDate, endDate).Scan(&res).Error; err != nil {
		return domain.Report{}, err
	}

	report.TotalRevenue = res.TotalRevenue
	report.TotalTransactions = res.TotalTransactions
	report.TopProduct.Name = res.TopProductName
	report.TopProduct.QtySold = res.TopProductQty

	return report, nil
}
