package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/notify-vital/backend/internal/auth"
	"github.com/notify-vital/backend/internal/models"
)

// AuthHandler provides methods to handle authentication routes
type AuthHandler struct {
	firebaseAuth   *auth.FirebaseAuth
	userRepository *auth.UserRepository
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(firebaseAuth *auth.FirebaseAuth, userRepository *auth.UserRepository) *AuthHandler {
	return &AuthHandler{
		firebaseAuth:   firebaseAuth,
		userRepository: userRepository,
	}
}

// SignUp handles user registration
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req models.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log request untuk debugging
	log.Printf("Attempting to create user with email: %s", req.Email)

	// Create user in Firebase Auth
	user, err := h.firebaseAuth.CreateUser(c.Request.Context(), &req)
	if err != nil {
		// Log the actual error from Firebase
		log.Printf("Firebase error: %v", err)

		if strings.Contains(err.Error(), "email already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	// Store additional user data in Firestore
	if err := h.userRepository.StoreUser(c.Request.Context(), user); err != nil {
		// If Firestore storage fails, still return success but log the error
		// In a production app, you would handle this more gracefully
		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully, but profile data storage failed",
			"uid":     user.UID,
		})
		return
	}

	// Create custom token
	token, err := h.firebaseAuth.CreateCustomToken(c.Request.Context(), user.UID)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully, but token generation failed",
			"uid":     user.UID,
		})
		return
	}

	c.JSON(http.StatusCreated, models.AuthResponse{
		Token:       token,
		ExpiresIn:   3600, // 1 hour
		DisplayName: user.DisplayName,
		UID:         user.UID,
		Email:       user.Email,
	})
}

// SignIn handles user login
func (h *AuthHandler) SignIn(c *gin.Context) {
	// Note: For Firebase Authentication, sign-in is typically handled on the client side
	// The server only verifies tokens and manages sessions
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign-in should be handled on the client side with Firebase Authentication SDK",
	})
}

// GetProfile handles getting user profile
func (h *AuthHandler) GetProfile(c *gin.Context) {
	uid := c.GetString("uid")
	if uid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get user from Firestore
	user, err := h.userRepository.GetUserByUID(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}

	if user == nil {
		// If user exists in Firebase Auth but not in Firestore, get basic info from Auth
		authUser, err := h.firebaseAuth.GetUserByUID(c.Request.Context(), uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
			return
		}

		// Store basic profile in Firestore for future
		authUser.CreatedAt = time.Now()
		if err := h.userRepository.StoreUser(c.Request.Context(), authUser); err != nil {
			// Just log the error, don't fail the request
		}

		user = authUser
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile handles updating user profile
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	uid := c.GetString("uid")
	if uid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update in Firebase Auth
	if err := h.firebaseAuth.UpdateUser(c.Request.Context(), uid, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user in Auth"})
		return
	}

	// Update in Firestore
	if err := h.userRepository.UpdateUser(c.Request.Context(), uid, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// DeleteAccount handles user account deletion
func (h *AuthHandler) DeleteAccount(c *gin.Context) {
	uid := c.GetString("uid")
	if uid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Delete from Firebase Auth
	if err := h.firebaseAuth.DeleteUser(c.Request.Context(), uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user account"})
		return
	}

	// Delete from Firestore
	if err := h.userRepository.DeleteUser(c.Request.Context(), uid); err != nil {
		// Just log the error, the auth deletion is more important
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}
