package domain

import "time"

// InvestmentStatus represents the status of an investment
type InvestmentStatus string

const (
	InvestmentStatusPending   InvestmentStatus = "pending"
	InvestmentStatusProcessed InvestmentStatus = "processed"
	InvestmentStatusCancelled InvestmentStatus = "cancelled"
)

// Investment represents a customer's investment in a fund
type Investment struct {
	ID         string           `json:"id"`
	CustomerID string           `json:"customer_id"`
	FundID     string           `json:"fund_id"`
	Amount     int64            `json:"amount"`
	Status     InvestmentStatus `json:"status"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}

// InvestmentRepository defines methods to interact with investments
type InvestmentRepository interface {
	GetByID(id string) (*Investment, error)
	GetByCustomerID(customerID string) ([]*Investment, error)
	Create(investment *Investment) error
	Update(investment *Investment) error
}

// InvestmentUseCase defines business logic for investments
type InvestmentUseCase interface {
	CreateInvestment(customerID, fundID string, amount int64) (*Investment, error)
	GetInvestment(id string) (*Investment, error)
	GetCustomerInvestments(customerID string) ([]*Investment, error)
}
