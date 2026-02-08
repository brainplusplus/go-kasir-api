package entity

import (
	"kasir-api/internal/domain"
	"time"
)

type Transaction struct {
	ID          int                 `gorm:"primaryKey"`
	TotalAmount int                 `gorm:"column:total_amount"`
	CreatedAt   time.Time           `gorm:"autoCreateTime"`
	Details     []TransactionDetail `gorm:"foreignKey:TransactionID"`
}

type TransactionDetail struct {
	ID            int    `gorm:"primaryKey"`
	TransactionID int    `gorm:"column:transaction_id"`
	ProductID     int    `gorm:"column:product_id"`
	ProductName   string `gorm:"column:product_name"`
	Quantity      int    `gorm:"column:qty"`
	Subtotal      int    `gorm:"column:subtotal"`
}

func (t *Transaction) ToDomain() domain.Transaction {
	details := make([]domain.TransactionDetail, 0)
	for _, d := range t.Details {
		details = append(details, d.ToDomain())
	}
	return domain.Transaction{
		ID:          t.ID,
		TotalAmount: t.TotalAmount,
		CreatedAt:   t.CreatedAt,
		Details:     details,
	}
}

func FromDomainTransaction(t domain.Transaction) Transaction {
	details := make([]TransactionDetail, 0)
	for _, d := range t.Details {
		details = append(details, FromDomainTransactionDetail(d))
	}
	return Transaction{
		ID:          t.ID,
		TotalAmount: t.TotalAmount,
		Details:     details,
	}
}

func (td *TransactionDetail) ToDomain() domain.TransactionDetail {
	return domain.TransactionDetail{
		ID:            td.ID,
		TransactionID: td.TransactionID,
		ProductID:     td.ProductID,
		ProductName:   td.ProductName,
		Quantity:      td.Quantity,
		Subtotal:      td.Subtotal,
	}
}

func FromDomainTransactionDetail(td domain.TransactionDetail) TransactionDetail {
	return TransactionDetail{
		ID:            td.ID,
		TransactionID: td.TransactionID,
		ProductID:     td.ProductID,
		ProductName:   td.ProductName,
		Quantity:      td.Quantity,
		Subtotal:      td.Subtotal,
	}
}
