package service

import "github.com/grokkos/go-isa-retail-service/internal/domain"

type fundService struct {
	fundRepo domain.FundRepository
}

// NewFundService creates a new instance of fund service
func NewFundService(fr domain.FundRepository) domain.FundService {
	return &fundService{
		fundRepo: fr,
	}
}

// GetFund gets a fund by ID
func (fs *fundService) GetFund(id string) (*domain.Fund, error) {
	return fs.fundRepo.GetByID(id)
}

// ListFunds lists all available funds
func (fs *fundService) ListFunds() ([]*domain.Fund, error) {
	return fs.fundRepo.GetAll()
}
