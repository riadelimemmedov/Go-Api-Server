package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logger "github.com/riadalimemmedov/api-server-project/api-server-project/pkg/log"
)

// ! GetOrders
func GetOrders(c *gin.Context) {
	orders := []map[string]interface{}{
		{"id": 1, "product": "Laptop", "quantity": 2, "price": 1200.00},
		{"id": 2, "product": "Mouse", "quantity": 5, "price": 25.00},
	}

	// Log the request
	logger.GetLogger().Info("Orders retrieved successfully")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Orders retrieved successfully",
		"data":    orders,
	})
}

// ! CreateOrder
func CreateOrder(c *gin.Context) {
	var order map[string]interface{}

	if err := c.ShouldBindJSON(&order); err != nil {
		logger.GetLogger().Info("Invalid request data")
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	order["id"] = 3
	order["created_at"] = "2025-06-04T10:00:00Z"

	// Log the request
	logger.GetLogger().Info("Order created successfully")

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Order created successfully",
		"data":    order,
	})
}


// ! GetOrderByID
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	order := map[string]interface{}{
		"id":       id,
		"product":  "Laptop",
		"quantity": 2,
		"price":    1200.00,
		"status":   "pending",
	}

	// Log the request
	logger.GetLogger().Info("Order return by ID")


	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Order retrieved successfully",
		"data":    order,
	})
}
