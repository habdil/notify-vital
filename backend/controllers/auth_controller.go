package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/habdil/notify-vital/backend/models"
	"github.com/habdil/notify-vital/backend/services"
)

// Register handles user registration
func Register(c *gin.Context) {
	// Parse request body
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register the user
	user, err := services.RegisterUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	// Generate token for the new user
	token, expiryTime, err := services.GenerateJWT(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token: " + err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, models.AuthResponse{
		Token:     token,
		ExpiresAt: expiryTime.Format(http.TimeFormat),
		User:      *user,
	})
}

// Login handles user authentication
func Login(c *gin.Context) {
	// Parse request body
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate the user
	user, token, expiryTime, err := services.LoginUser(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to authenticate: " + err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, models.AuthResponse{
		Token:     token,
		ExpiresAt: expiryTime.Format(http.TimeFormat),
		User:      *user,
	})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// Get token from context (set by auth middleware)
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token found"})
		return
	}

	// Invalidate the token
	err := services.LogoutUser(token.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout: " + err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// Me retrieves the authenticated user's profile
func Me(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	// Get user details
	user, err := services.GetUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data: " + err.Error()})
		return
	}

	// Return user data
	c.JSON(http.StatusOK, gin.H{"user": user})
}
