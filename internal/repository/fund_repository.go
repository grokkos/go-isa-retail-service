package repository

import (
	"errors"
	"github.com/grokkos/go-isa-retail-service/internal/domain"
	"sync"
	"time"
)

type inMemoryFundRepository struct {
	mutex sync.RWMutex
	funds map[string]*domain.Fund
}

// NewInMemoryFundRepository creates a new in-memory fund repository
func NewInMemoryFundRepository() domain.FundRepository {
	// Initialize with sample funds
	funds := map[string]*domain.Fund{
		"fund-1": {
			ID:          "fund-1",
			Name:        "Cushon Equities Fund",
			Description: "A fund that invests in global equities for long-term growth",
			RiskLevel:   domain.RiskLevelHigh,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		"fund-2": {
			ID:          "fund-2",
			Name:        "Cushon Balanced Fund",
			Description: "A balanced fund that invests in a mix of equities and bonds",
			RiskLevel:   domain.RiskLevelMedium,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		"fund-3": {
			ID:          "fund-3",
			Name:        "Cushon Bond Fund",
			Description: "A fund that invests in government and corporate bonds",
			RiskLevel:   domain.RiskLevelLow,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return &inMemoryFundRepository{
		funds: funds,
	}
}

// GetByID gets a fund by ID
func (r *inMemoryFundRepository) GetByID(id string) (*domain.Fund, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	fund, ok := r.funds[id]
	if !ok {
		return nil, errors.New("fund not found")
	}

	return fund, nil
}

// GetAll gets all funds
func (r *inMemoryFundRepository) GetAll() ([]*domain.Fund, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	funds := make([]*domain.Fund, 0, len(r.funds))
	for _, fund := range r.funds {
		funds = append(funds, fund)
	}

	return funds, nil
}
