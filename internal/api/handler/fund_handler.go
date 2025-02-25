package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/grokkos/go-isa-retail-service/internal/domain"
	"net/http"
)

// FundHandler handles HTTP requests related to funds
type FundHandler struct {
	FundUseCase domain.FundService
}

// NewFundHandler creates a new fund handler
func NewFundHandler(fu domain.FundService) *FundHandler {
	return &FundHandler{
		FundUseCase: fu,
	}
}

// GetFund handles GET /funds/{id}
func (h *FundHandler) GetFund(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fund, err := h.FundUseCase.GetFund(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fund)
}

// ListFunds handles GET /funds
func (h *FundHandler) ListFunds(w http.ResponseWriter, r *http.Request) {
	funds, err := h.FundUseCase.ListFunds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(funds)
}
