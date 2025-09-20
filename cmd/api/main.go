package main

import (
	"github.com/Ch94Ca/ms-nexusMarket-inventory/internal/app/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", handler.HealthCheckHandler)
	error := r.Run(":8090")
	if error != nil {
		panic(error)
	}
}
