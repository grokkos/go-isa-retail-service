package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/grokkos/go-isa-retail-service/internal/domain"
	"net/http"
	"strconv"
)

// InvestmentHandler handles HTTP requests related to investments
type InvestmentHandler struct {
	InvestmentService domain.InvestmentService
	FundService       domain.FundService
}

// NewInvestmentHandler creates a new investment handler
func NewInvestmentHandler(is domain.InvestmentService, fs domain.FundService) *InvestmentHandler {
	return &InvestmentHandler{
		InvestmentService: is,
		FundService:       fs,
	}
}

// CreateInvestmentRequest is the request for creating an investment
type CreateInvestmentRequest struct {
	CustomerID string `json:"customer_id"`
	FundID     string `json:"fund_id"`
	Amount     string `json:"amount"` // Amount as string (e.g., "25000.00")
}

// CreateInvestmentResponse is the response for creating an investment
type CreateInvestmentResponse struct {
	ID          string  `json:"id"`
	CustomerID  string  `json:"customer_id"`
	FundID      string  `json:"fund_id"`
	FundName    string  `json:"fund_name"`
	Amount      string  `json:"amount"`
	AmountValue float64 `json:"amount_value"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
}

// CreateInvestment handles POST /investments
func (h *InvestmentHandler) CreateInvestment(w http.ResponseWriter, r *http.Request) {
	var req CreateInvestmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert amount to pence (int64)
	amountFloat, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		http.Error(w, "Invalid amount format", http.StatusBadRequest)
		return
	}
	amountPence := int64(amountFloat * 100)

	investment, err := h.InvestmentService.CreateInvestment(req.CustomerID, req.FundID, amountPence)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Properly look up the fund name using the FundService
	var fundName string
	fund, err := h.FundService.GetFund(investment.FundID)
	if err != nil {
		// If we can't find the fund, use a placeholder but don't fail the request
		fundName = "Unknown Fund"
	} else {
		fundName = fund.Name
	}

	response := CreateInvestmentResponse{
		ID:          investment.ID,
		CustomerID:  investment.CustomerID,
		FundID:      investment.FundID,
		FundName:    fundName,
		Amount:      req.Amount,
		AmountValue: float64(investment.Amount) / 100.0,
		Status:      string(investment.Status),
		CreatedAt:   investment.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetInvestment handles GET /investments/{id}
func (h *InvestmentHandler) GetInvestment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	investment, err := h.InvestmentService.GetInvestment(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Enrich the response with fund information
	fund, err := h.FundService.GetFund(investment.FundID)
	var fundName string
	if err != nil {
		fundName = "Unknown Fund"
	} else {
		fundName = fund.Name
	}

	// Create a more detailed response
	response := struct {
		ID         string  `json:"id"`
		CustomerID string  `json:"customer_id"`
		FundID     string  `json:"fund_id"`
		FundName   string  `json:"fund_name"`
		Amount     float64 `json:"amount"`
		Status     string  `json:"status"`
		CreatedAt  string  `json:"created_at"`
		UpdatedAt  string  `json:"updated_at"`
	}{
		ID:         investment.ID,
		CustomerID: investment.CustomerID,
		FundID:     investment.FundID,
		FundName:   fundName,
		Amount:     float64(investment.Amount) / 100.0,
		Status:     string(investment.Status),
		CreatedAt:  investment.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  investment.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCustomerInvestments handles GET /customers/{id}/investments
func (h *InvestmentHandler) GetCustomerInvestments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]

	investments, err := h.InvestmentService.GetCustomerInvestments(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enrich the responses with fund information
	type EnrichedInvestment struct {
		ID         string  `json:"id"`
		CustomerID string  `json:"customer_id"`
		FundID     string  `json:"fund_id"`
		FundName   string  `json:"fund_name"`
		Amount     float64 `json:"amount"`
		Status     string  `json:"status"`
		CreatedAt  string  `json:"created_at"`
	}

	enrichedInvestments := make([]EnrichedInvestment, 0, len(investments))
	for _, investment := range investments {
		// Get fund name from FundService
		fund, err := h.FundService.GetFund(investment.FundID)
		fundName := "Unknown Fund"
		if err == nil {
			fundName = fund.Name
		}

		enriched := EnrichedInvestment{
			ID:         investment.ID,
			CustomerID: investment.CustomerID,
			FundID:     investment.FundID,
			FundName:   fundName,
			Amount:     float64(investment.Amount) / 100.0,
			Status:     string(investment.Status),
			CreatedAt:  investment.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		enrichedInvestments = append(enrichedInvestments, enriched)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(enrichedInvestments)
}
