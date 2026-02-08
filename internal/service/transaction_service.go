package service

import (
	"context"
	"errors"
	"kasir-api/internal/adapter/inbound/http/dto"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/port/outbound"
	"time"
)

type transactionService struct {
	repo       outbound.TransactionRepository
	produkRepo outbound.ProdukRepository
	reportRepo outbound.ReportRepository
	tm         outbound.TransactionManager
}

func NewTransactionService(repo outbound.TransactionRepository, produkRepo outbound.ProdukRepository, reportRepo outbound.ReportRepository, tm outbound.TransactionManager) inbound.TransactionService {
	return &transactionService{repo: repo, produkRepo: produkRepo, reportRepo: reportRepo, tm: tm}
}

func (s *transactionService) Checkout(ctx context.Context, request dto.CheckoutRequest) (domain.Transaction, error) {
	var transaction domain.Transaction

	err := s.tm.RunInTransaction(ctx, func(ctx context.Context) error {
		var details []domain.TransactionDetail
		var totalAmount int

		for _, item := range request.Items {
			// Get Product Data (Price, Name, Stock)
			product, err := s.produkRepo.FindByID(ctx, item.ProductID)
			if err != nil {
				return err
			}

			// Validate Stock (Business Logic)
			if product.Stock < item.Quantity {
				return errors.New("insufficient stock for product: " + product.Name)
			}

			// Deduct Stock
			if err := s.produkRepo.UpdateStock(ctx, item.ProductID, item.Quantity); err != nil {
				return err
			}

			// Calculate Subtotal (Business Logic)
			subtotal := product.Price * item.Quantity
			totalAmount += subtotal

			details = append(details, domain.TransactionDetail{
				ProductID:   product.ID,
				ProductName: product.Name,
				Quantity:    item.Quantity,
				Subtotal:    subtotal,
			})
		}

		transaction = domain.Transaction{
			TotalAmount: totalAmount,
			Details:     details,
		}

		// Persist Transaction
		createdTransaction, err := s.repo.Create(ctx, transaction)
		if err != nil {
			return err
		}
		transaction = createdTransaction
		return nil
	})

	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (s *transactionService) GetReport(ctx context.Context, startDate, endDate time.Time) (domain.Report, error) {
	// Use ReportRepository for queries
	return s.reportRepo.GetStatsByRange(ctx, startDate, endDate)
}

func (s *transactionService) GetReportToday(ctx context.Context) (domain.Report, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	return s.reportRepo.GetStatsByRange(ctx, startOfDay, endOfDay)
}
