package repository

import (
	"errors"
	"github.com/grokkos/go-isa-retail-service/internal/domain"
	"sync"
)

type inMemoryInvestmentRepository struct {
	mutex       sync.RWMutex
	investments map[string]*domain.Investment
}

// NewInMemoryInvestmentRepository creates a new in-memory investment repository
func NewInMemoryInvestmentRepository() domain.InvestmentRepository {
	return &inMemoryInvestmentRepository{
		investments: make(map[string]*domain.Investment),
	}
}

// GetByID gets an investment by ID
func (r *inMemoryInvestmentRepository) GetByID(id string) (*domain.Investment, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	investment, ok := r.investments[id]
	if !ok {
		return nil, errors.New("investment not found")
	}

	return investment, nil
}

// GetByCustomerID gets all investments for a customer
func (r *inMemoryInvestmentRepository) GetByCustomerID(customerID string) ([]*domain.Investment, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var investments []*domain.Investment
	for _, investment := range r.investments {
		if investment.CustomerID == customerID {
			investments = append(investments, investment)
		}
	}

	return investments, nil
}

// Create creates a new investment
func (r *inMemoryInvestmentRepository) Create(investment *domain.Investment) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.investments[investment.ID]; ok {
		return errors.New("investment already exists")
	}

	r.investments[investment.ID] = investment
	return nil
}

// Update updates an existing investment
func (r *inMemoryInvestmentRepository) Update(investment *domain.Investment) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.investments[investment.ID]; !ok {
		return errors.New("investment not found")
	}

	r.investments[investment.ID] = investment
	return nil
}
