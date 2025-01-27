package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp handles user signup
func SignUp(c *gin.Context) {
	// Struct for request payload
	type SignUpRequest struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	// Parse and validate JSON request body
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Logic to save user data (placeholder for actual database logic)
	// In real-world apps, you would hash the password and save the user in a database.
	// For example:
	// hashedPassword := hashPassword(req.Password)
	// Save to database logic goes here...

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message":  "User successfully signed up!",
		"username": req.Username,
		"email":    req.Email,
	})
}
