package api

import "cash-flow-service/models"

type Service interface {
	GetCashFlows() ([]models.CashFlow, error)
	AddCashFlow(cashFlow models.CashFlow) error
	UpdateCashFlow(cashFlow models.CashFlow) error
}
