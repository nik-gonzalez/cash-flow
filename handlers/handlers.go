package handlers

import (
	"cash-flow-service/api"
	"cash-flow-service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IHandler interface {
	GetCashFlows(ctx *gin.Context)
	AddCashFlow(ctx *gin.Context)
}

type Handler struct {
	service api.Service
}

func NewWithService(service api.Service) Handler {
	return Handler{service: service}
}
func (handler Handler) GetCashFlows(c *gin.Context) {
	cashFlows, err := handler.service.GetCashFlows()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	} else {
		fmt.Println("handling GetCashFlows", cashFlows)
		c.IndentedJSON(http.StatusOK, cashFlows)
	}
}

func (handler Handler) AddCashFlows(c *gin.Context) {
	var cashFlow models.CashFlow
	if err := c.BindJSON(&cashFlow); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	} else {
		fmt.Println("handling AddCashFlows", cashFlow)
		err := handler.service.AddCashFlow(cashFlow)
		c.Status(http.StatusCreated)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
}

func (handler Handler) UpdateCashFlows(c *gin.Context) {
	var cashFlow models.CashFlow
	if err := c.BindJSON(&cashFlow); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	} else {
		fmt.Println("handling UpdateCashFlows", cashFlow)
		err := handler.service.UpdateCashFlow(cashFlow)
		c.Status(http.StatusOK)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}
}
