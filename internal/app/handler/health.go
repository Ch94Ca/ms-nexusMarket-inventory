// Package handler implements HTTP handlers for the application.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler handles the health check endpoint.
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "stock-service up",
	})
}
