package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Owner  string `json:"owner" binding:"required" `
	Status string `json:"status" binding:"required,status"`
}

// GetExample handles the GET request for the example endpoint
func GetExample(c *gin.Context) {
	// Handle your logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "GET Example",
	})
}

// CreateExample handles the POST request for the example endpoint
func CreateExample(c *gin.Context) {
	// Handle your logic here
	c.JSON(http.StatusCreated, gin.H{
		"message": "Create Example",
	})
}

// UpdateExample handles the PUT request for the example endpoint
func UpdateExample(c *gin.Context) {
	// Handle your logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Example",
	})
}

// DeleteExample handles the DELETE request for the example endpoint
func DeleteExample(c *gin.Context) {
	// Handle your logic here
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Example",
	})
}
