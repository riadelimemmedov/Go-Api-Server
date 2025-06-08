package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ! GetUsers return list a of the users
func GetUsers(c *gin.Context) {
	users := []map[string]interface{}{
		{"id": 1, "name": "John Doe", "email": "john@example.com"},
		{"id": 2, "name": "Jane Smith", "email": "jane@example.com"},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

// ! CreateUser and save to db
func CreateUser(c *gin.Context) {
	var user map[string]interface{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	// Add ID and timestamp
	user["id"] = 3
	user["created_at"] = "2025-06-04T10:00:00Z"

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
		"data":    user,
	})
}

// !GetUserByID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user := map[string]interface{}{
		"id":     id,
		"name":   "John Doe",
		"email":  "john@example.com",
		"status": "active",
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User retrieved successfully",
		"data":    user,
	})
}
