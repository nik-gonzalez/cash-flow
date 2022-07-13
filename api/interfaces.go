package api

import (
	"cash-flow-service/dto"
	"cash-flow-service/model"
	"github.com/gin-gonic/gin"
)

type Repo interface {
	Create(cashFlow model.CashFlowEntity) (model.CashFlowEntity, error)
	Update(cashFlow model.CashFlowEntity) (model.CashFlowEntity, error)
	GetById(id int) (model.CashFlowEntity, error)
	GetAllBySearchCriteria(command CashFlowSearchCriteria)
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
