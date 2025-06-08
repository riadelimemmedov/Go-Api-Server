package router

import (
	"github.com/gin-gonic/gin"
	"github.com/riadalimemmedov/api-server-project/api-server-project/internal/handlers"
	"github.com/riadalimemmedov/api-server-project/api-server-project/internal/middleware"
)

// ! SetupRouter
func SetupRouter() *gin.Engine {
	r := gin.New()

	// Middleware
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// Health check
	r.GET("/health", handlers.HealthCheck)

	// API routes
	api := r.Group("api/v1")
	{
		//Orders
		api.GET("/orders", handlers.GetOrders)
		api.POST("/orders", handlers.CreateOrder)
		api.GET("/orders/:id", handlers.GetOrderByID)

		// Users
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)
		api.GET("/users/:id", handlers.GetUserByID)
	}
	return r
}
