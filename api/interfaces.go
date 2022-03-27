package api

import (
	"cash-flow-service/models"
	"github.com/gin-gonic/gin"
)

type IService interface {
	GetCashFlows() ([]models.CashFlow, error)
	AddCashFlow(cashFlow models.CashFlow) error
	UpdateCashFlow(cashFlow models.CashFlow) error
	GetCashFlow(id string) (*models.CashFlow, error)
	DeleteCashFlow(id string) error
}

type IHandler interface {
	GetCashFlows(ctx *gin.Context)
	AddCashFlow(c *gin.Context)
	UpdateCashFlows(c *gin.Context)
	GetCashFlow(c *gin.Context)
	DeleteCashFlow(c *gin.Context)
}
