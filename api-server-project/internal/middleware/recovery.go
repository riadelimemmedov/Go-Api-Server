package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logger "github.com/riadalimemmedov/api-server-project/api-server-project/pkg/log"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.GetLogger().Info("Panic recovered")

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal server error",
			"error":   "Something went wrong",
		})
	})
}
