package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ! HealthCheck return status of the server
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"message":   "Service is healthy",
		"timestamp": "2025-06-04T10:00:00Z",
	})
}
