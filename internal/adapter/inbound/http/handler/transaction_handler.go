package handler

import (
	"encoding/json"
	"kasir-api/internal/adapter/inbound/http/dto"
	"kasir-api/internal/domain"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/utils/swagger"
	"net/http"
	"time"
)

type TransactionHandler struct {
	service inbound.TransactionService
	swag    *swagger.Generator
}

func NewTransactionHandler(service inbound.TransactionService, swag *swagger.Generator) *TransactionHandler {
	return &TransactionHandler{service: service, swag: swag}
}

func (h *TransactionHandler) RegisterRoutes(mux *http.ServeMux) {
	// Register to Swagger
	h.swag.Register("POST", "/api/checkout", "Checkout Transaction", dto.CheckoutRequest{}, dto.TransactionResponse{})
	h.swag.Register("GET", "/api/report/hari-ini", "Get Daily Report", nil, dto.ReportResponse{})
	h.swag.RegisterWithQuery("GET", "/api/report", "Get Report (Range)", nil, dto.ReportResponse{}, []swagger.QueryParam{
		{Name: "start_date", Type: "string", Description: "Start Date (YYYY-MM-DD)", Required: false},
		{Name: "end_date", Type: "string", Description: "End Date (YYYY-MM-DD)", Required: false},
	})

	// Register to Mux
	mux.HandleFunc("/api/checkout", h.handleCheckout)
	mux.HandleFunc("/api/report/hari-ini", h.handleReport)
	mux.HandleFunc("/api/report", h.handleReport)
}

func (h *TransactionHandler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dto.CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(r.Context(), req)
	if err != nil {
		// Distinction between validation error (stock) and server error could be improved
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := dto.FromDomainTransaction(transaction)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *TransactionHandler) handleReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	var report domain.Report
	var err error

	if startDateStr != "" && endDateStr != "" {
		var startDate, endDate time.Time
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			http.Error(w, "Invalid start_date format (YYYY-MM-DD)", http.StatusBadRequest)
			return
		}
		endDateParse, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			http.Error(w, "Invalid end_date format (YYYY-MM-DD)", http.StatusBadRequest)
			return
		}
		// End date should be exclusive or end of day? Usually inclusive for "2026-02-01". Be inclusive.
		// Let's make it up to end of that day.
		endDate = endDateParse.Add(24 * time.Hour)

		report, err = h.service.GetReport(r.Context(), startDate, endDate)
	} else {
		// Default to Today
		report, err = h.service.GetReportToday(r.Context())
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.FromDomainReport(report)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
