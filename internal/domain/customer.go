package domain

import "time"

// Customer represents a retail customer who can make ISA investments
type Customer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomerRepository defines methods to interact with customers
type CustomerRepository interface {
	GetByID(id string) (*Customer, error)
	Create(customer *Customer) error
	Update(customer *Customer) error
}
