package main

import (
	"cash-flow-service/handlers"
	"cash-flow-service/memory"
	"github.com/gin-gonic/gin"
)

func main() {
	service := memory.InMemoryService{}
	handler := handlers.NewWithService(service)
	router := gin.Default()
	router.GET("/cash-flows", handler.GetCashFlows)
	router.POST("/cash-flows", handler.AddCashFlows)
	router.PUT("/cash-flows", handler.UpdateCashFlows)
	router.Run("localhost:8080")
}
