package memory

import (
	"cash-flow-service/models"
	"errors"
	"time"
)

var cashFlows = map[string]models.CashFlow{
	"1": {ID: "1", Name: "Salary", Remarks: "Monthly Salary from Main Gig", Amount: 39.99, Date: models.JsonDateTime{Time: time.Now()}},
	"2": {ID: "2", Name: "Side Hustle Income", Remarks: "Project B", Amount: 56.99, Date: models.JsonDateTime{Time: time.Now()}},
	"3": {ID: "3", Name: "Food", Remarks: "Grocery purchase", Amount: -17.99, Date: models.JsonDateTime{Time: time.Now()}},
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

func (service InMemoryService) GetCashFlow(id string) (*models.CashFlow, error) {
	cashFlow, exists := cashFlows[id]
	if !exists {
		return nil, errors.New("CashFlow not found")
	}
	return &cashFlow, nil
}

func (service InMemoryService) DeleteCashFlow(id string) error {
	_, exists := cashFlows[id]
	if !exists {
		return errors.New("CashFlow not found")
	}
	delete(cashFlows, id)
	return nil
}
