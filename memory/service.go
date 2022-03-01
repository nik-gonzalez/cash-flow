package memory

import (
	"cash-flow-service/models"
	"errors"
)

// albums slice to seed record album data.
// albums slice to seed record album data.
var cashFlows = map[string]models.CashFlow{
	"1": {ID: "1", Title: "Side Hustle Income", Description: "John Coltrane", Amount: 56.99},
	"2": {ID: "2", Title: "Food", Description: "Gerry Mulligan", Amount: -17.99},
	"3": {ID: "3", Title: "Salary", Description: "Sarah Vaughan", Amount: 39.99},
}

type InMemoryService struct{}

func (service InMemoryService) GetCashFlows() ([]models.CashFlow, error) {
	l := make([]models.CashFlow, 0, len(cashFlows))
	for _, value := range cashFlows {
		l = append(l, value)
	}
	return l, nil
}

func (service InMemoryService) AddCashFlow(cashFlow models.CashFlow) error {
	if _, exists := cashFlows[cashFlow.ID]; exists {
		return errors.New("CashFlow already exists")
	}
	cashFlows[cashFlow.ID] = cashFlow
	return nil
}

func (service InMemoryService) UpdateCashFlow(cashFlow models.CashFlow) error {
	if _, exists := cashFlows[cashFlow.ID]; !exists {
		return errors.New("CashFlow not found")
	}
	cashFlows[cashFlow.ID] = cashFlow
	return nil
}
