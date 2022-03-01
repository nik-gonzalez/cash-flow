package api

import "cash-flow-service/models"

type Service interface {
	GetCashFlows() ([]models.CashFlow, error)
	AddCashFlow(cashFlow models.CashFlow) error
	UpdateCashFlow(cashFlow models.CashFlow) error
	GetCashFlow(id string) (*models.CashFlow, error)
	DeleteCashFlow(id string) error
}
