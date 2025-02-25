package repository

import (
	"errors"
	"github.com/grokkos/go-isa-retail-service/internal/domain"
	"sync"
	"time"
)

type inMemoryCustomerRepository struct {
	mutex     sync.RWMutex
	customers map[string]*domain.Customer
}

// NewInMemoryCustomerRepository creates a new in-memory customer repository
func NewInMemoryCustomerRepository() domain.CustomerRepository {
	// Initialize with sample customer
	customers := map[string]*domain.Customer{
		"customer-1": {
			ID:        "customer-1",
			Name:      "John Smith",
			Email:     "john.smith@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return &inMemoryCustomerRepository{
		customers: customers,
	}
}

// GetByID gets a customer by ID
func (r *inMemoryCustomerRepository) GetByID(id string) (*domain.Customer, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	customer, ok := r.customers[id]
	if !ok {
		return nil, errors.New("customer not found")
	}

	return customer, nil
}

// Create creates a new customer
func (r *inMemoryCustomerRepository) Create(customer *domain.Customer) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.customers[customer.ID]; ok {
		return errors.New("customer already exists")
	}

	r.customers[customer.ID] = customer
	return nil
}

// Update updates an existing customer
func (r *inMemoryCustomerRepository) Update(customer *domain.Customer) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.customers[customer.ID]; !ok {
		return errors.New("customer not found")
	}

	r.customers[customer.ID] = customer
	return nil
}
