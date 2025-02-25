package service_test

import (
	"github.com/grokkos/go-isa-retail-service/internal/domain"
	"github.com/grokkos/go-isa-retail-service/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Mock InvestmentRepository
type mockInvestmentRepository struct {
	mock.Mock
}

func (m *mockInvestmentRepository) GetByID(id string) (*domain.Investment, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Investment), args.Error(1)
}

func (m *mockInvestmentRepository) GetByCustomerID(customerID string) ([]*domain.Investment, error) {
	args := m.Called(customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Investment), args.Error(1)
}

func (m *mockInvestmentRepository) Create(investment *domain.Investment) error {
	args := m.Called(investment)
	return args.Error(0)
}

func (m *mockInvestmentRepository) Update(investment *domain.Investment) error {
	args := m.Called(investment)
	return args.Error(0)
}

// Mock CustomerRepository
type mockCustomerRepository struct {
	mock.Mock
}

func (m *mockCustomerRepository) GetByID(id string) (*domain.Customer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Customer), args.Error(1)
}

func (m *mockCustomerRepository) Create(customer *domain.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *mockCustomerRepository) Update(customer *domain.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

// Mock FundRepository
type mockFundRepository struct {
	mock.Mock
}

func (m *mockFundRepository) GetByID(id string) (*domain.Fund, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Fund), args.Error(1)
}

func (m *mockFundRepository) GetAll() ([]*domain.Fund, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Fund), args.Error(1)
}

func TestInvestmentValidation(t *testing.T) {
	// Create separate mocks for each repository type
	mockInvestRepo := new(mockInvestmentRepository)
	mockCustomerRepo := new(mockCustomerRepository)
	mockFundRepo := new(mockFundRepository)

	// Set up the investment service
	investmentService := service.NewInvestmentService(
		mockInvestRepo,
		mockCustomerRepo,
		mockFundRepo,
	)

	// Set up test data
	mockCustomer := &domain.Customer{ID: "customer-1", Name: "Test Customer"}
	mockFund := &domain.Fund{ID: "fund-1", Name: "Test Fund"}

	// Configure mocks to return our test data
	mockCustomerRepo.On("GetByID", "customer-1").Return(mockCustomer, nil)
	mockFundRepo.On("GetByID", "fund-1").Return(mockFund, nil)
	mockInvestRepo.On("Create", mock.AnythingOfType("*domain.Investment")).Return(nil)

	// Test case 1: Successful investment within ISA limit
	t.Run("Valid investment within ISA limit", func(t *testing.T) {
		investment, err := investmentService.CreateInvestment("customer-1", "fund-1", 1500000) // £15,000
		assert.NoError(t, err)
		assert.NotNil(t, investment)
		assert.Equal(t, int64(1500000), investment.Amount)
	})

	// Test case 2: Investment exceeding ISA limit
	t.Run("Investment exceeding ISA limit", func(t *testing.T) {
		investment, err := investmentService.CreateInvestment("customer-1", "fund-1", 2500000) // £25,000
		assert.Error(t, err)
		assert.Nil(t, investment)
		assert.Contains(t, err.Error(), "exceeds ISA annual limit")
	})
}
