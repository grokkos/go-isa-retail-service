package domain

import "time"

// RiskLevel represents the risk level of a fund
type RiskLevel string

const (
	RiskLevelLow    RiskLevel = "low"
	RiskLevelMedium RiskLevel = "medium"
	RiskLevelHigh   RiskLevel = "high"
)

// Fund represents an investment fund that customers can invest in
type Fund struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	RiskLevel   RiskLevel `json:"risk_level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// FundRepository defines methods to interact with funds
type FundRepository interface {
	GetByID(id string) (*Fund, error)
	GetAll() ([]*Fund, error)
}

// FundService defines business logic for funds
type FundService interface {
	GetFund(id string) (*Fund, error)
	ListFunds() ([]*Fund, error)
}
