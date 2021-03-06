package main

import (
	"cash-flow-service/handlers"
	"cash-flow-service/memory"
	"github.com/gin-gonic/gin"
)

func main() {
	service := memory.InMemoryService{}
	handler := handlers.NewHandler(service)
	router := gin.Default()
	router.GET("/cash-flows", handler.GetCashFlows)
	router.GET("/cash-flows/:id", handler.GetCashFlow)
	router.POST("/cash-flows", handler.AddCashFlow)
	router.PUT("/cash-flows", handler.UpdateCashFlow)
	router.DELETE("/cash-flows/:id", handler.DeleteCashFlow)
	router.Run("localhost:8080")
}
