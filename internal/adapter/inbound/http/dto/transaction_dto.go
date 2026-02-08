package dto

import (
	"kasir-api/internal/domain"
	"time"
)

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"qty"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type TransactionDetailResponse struct {
	ID          int    `json:"id"`
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"qty"`
	Subtotal    int    `json:"subtotal"`
}

type TransactionResponse struct {
	ID          int                         `json:"id"`
	TotalAmount int                         `json:"total_amount"`
	CreatedAt   time.Time                   `json:"created_at"`
	Details     []TransactionDetailResponse `json:"details"`
}

type ReportTopProductResponse struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type ReportResponse struct {
	TotalRevenue   int                      `json:"total_revenue"`
	TotalTransaksi int                      `json:"total_transaksi"`
	ProdukTerlaris ReportTopProductResponse `json:"produk_terlaris"`
}

func FromDomainTransaction(t domain.Transaction) TransactionResponse {
	details := make([]TransactionDetailResponse, 0)
	for _, d := range t.Details {
		details = append(details, TransactionDetailResponse{
			ID:          d.ID,
			ProductID:   d.ProductID,
			ProductName: d.ProductName,
			Quantity:    d.Quantity,
			Subtotal:    d.Subtotal,
		})
	}
	return TransactionResponse{
		ID:          t.ID,
		TotalAmount: t.TotalAmount,
		CreatedAt:   t.CreatedAt,
		Details:     details,
	}
}

func FromDomainReport(r domain.Report) ReportResponse {
	return ReportResponse{
		TotalRevenue:   r.TotalRevenue,
		TotalTransaksi: r.TotalTransactions,
		ProdukTerlaris: ReportTopProductResponse{
			Nama:       r.TopProduct.Name,
			QtyTerjual: r.TopProduct.QtySold,
		},
	}
}
