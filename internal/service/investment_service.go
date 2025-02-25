package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/grokkos/go-isa-retail-service/internal/domain"
	"time"
)

type investmentService struct {
	investmentRepo domain.InvestmentRepository
	customerRepo   domain.CustomerRepository
	fundRepo       domain.FundRepository
}

// NewInvestmentService creates a new instance of investment service
func NewInvestmentService(
	ir domain.InvestmentRepository,
	cr domain.CustomerRepository,
	fr domain.FundRepository,
) domain.InvestmentService {
	return &investmentService{
		investmentRepo: ir,
		customerRepo:   cr,
		fundRepo:       fr,
	}
}

// CreateInvestment creates a new investment
func (is *investmentService) CreateInvestment(customerID, fundID string, amount int64) (*domain.Investment, error) {
	// Check if customer exists
	customer, err := is.customerRepo.GetByID(customerID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}

	// Check if fund exists
	fund, err := is.fundRepo.GetByID(fundID)
	if err != nil {
		return nil, err
	}
	if fund == nil {
		return nil, errors.New("fund not found")
	}

	// Validate amount
	if amount <= 0 {
		return nil, errors.New("investment amount must be positive")
	}

	// ISA annual limit check - this would be more comprehensive in a real implementation
	const ISA_ANNUAL_LIMIT_IN_PENCE = 2000000 // £20,000 in pence
	if amount > ISA_ANNUAL_LIMIT_IN_PENCE {
		return nil, errors.New("investment exceeds ISA annual limit of £20,000")
	}

	// Create investment
	investment := &domain.Investment{
		ID:         uuid.New().String(),
		CustomerID: customerID,
		FundID:     fundID,
		Amount:     amount,
		Status:     domain.InvestmentStatusPending,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Save investment
	err = is.investmentRepo.Create(investment)
	if err != nil {
		return nil, err
	}

	return investment, nil
}

// GetInvestment gets an investment by ID
func (is *investmentService) GetInvestment(id string) (*domain.Investment, error) {
	return is.investmentRepo.GetByID(id)
}

// GetCustomerInvestments gets all investments for a customer
func (is *investmentService) GetCustomerInvestments(customerID string) ([]*domain.Investment, error) {
	return is.investmentRepo.GetByCustomerID(customerID)
}
