package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/notify-vital/backend/internal/auth"
)

// AuthMiddleware creates a middleware that validates Firebase ID tokens
func AuthMiddleware(firebaseAuth *auth.FirebaseAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		// Format: "Bearer <token>"
		idToken := ""
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) == 2 {
			idToken = strings.TrimSpace(parts[1])
		}

		if idToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ID token is required"})
			c.Abort()
			return
		}

		// Verify the ID token
		token, err := firebaseAuth.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			c.Abort()
			return
		}

		// Store the verified UID in the context
		c.Set("uid", token.UID)
		c.Next()
	}
}

// OptionalAuthMiddleware creates a middleware that attempts to validate Firebase ID tokens,
// but doesn't abort if token is invalid or missing
func OptionalAuthMiddleware(firebaseAuth *auth.FirebaseAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// Extract the token from the Authorization header
		idToken := ""
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) == 2 {
			idToken = strings.TrimSpace(parts[1])
		}

		if idToken == "" {
			c.Next()
			return
		}

		// Verify the ID token
		token, err := firebaseAuth.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			// Just continue without setting the UID
			c.Next()
			return
		}

		// Store the verified UID in the context
		c.Set("uid", token.UID)
		c.Next()
	}
}
