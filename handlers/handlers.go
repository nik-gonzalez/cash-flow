package handlers

import (
	"cash-flow-service/api"
	"cash-flow-service/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	service api.IService
}

func NewHandler(service api.IService) *Handler {
	return &Handler{service}
}

func (handler Handler) GetCashFlows(c *gin.Context) {
	cashFlows, err := handler.service.GetCashFlows()
	logrus.Println("handling GetCashFlows")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, cashFlows)
}

func (handler Handler) AddCashFlow(c *gin.Context) {
	var cashFlow models.CashFlow
	if err := c.BindJSON(&cashFlow); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	logrus.Println("handling AddCashFlow", cashFlow)
	err := handler.service.AddCashFlow(cashFlow)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (handler Handler) UpdateCashFlow(c *gin.Context) {
	var cashFlow models.CashFlow
	if err := c.BindJSON(&cashFlow); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	logrus.Println("handling UpdateCashFlow", cashFlow)
	err := handler.service.UpdateCashFlow(cashFlow)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (handler Handler) GetCashFlow(c *gin.Context) {
	var cashFlow *models.CashFlow
	var err error
	id := c.Param("id")
	logrus.Println("handling GetCashFlow", id)
	cashFlow, err = handler.service.GetCashFlow(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, cashFlow)
}

func (handler Handler) DeleteCashFlow(c *gin.Context) {
	id := c.Param("id")
	logrus.Println("handling DeleteCashFlow", id)
	err := handler.service.DeleteCashFlow(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
