package api

import (
	"cash-flow-service/dto"
	"github.com/gin-gonic/gin"
)

type Repo interface {
	Save()
}

type IService interface {
	GetCashFlows() ([]dto.CashFlow, error)
	AddCashFlow(cashFlow dto.CashFlow) error
	UpdateCashFlow(cashFlow dto.CashFlow) error
	GetCashFlow(id string) (*dto.CashFlow, error)
	DeleteCashFlow(id string) error
}

type IHandler interface {
	GetCashFlows(ctx *gin.Context)
	AddCashFlow(c *gin.Context)
	UpdateCashFlow(c *gin.Context)
	GetCashFlow(c *gin.Context)
	DeleteCashFlow(c *gin.Context)
}
