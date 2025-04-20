package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/habdil/notify-vital/backend/controllers"
	"github.com/habdil/notify-vital/backend/middleware"
)

// SetupAuthRoutes configures all authentication routes
func SetupAuthRoutes(router *gin.Engine) {
	// Public routes (no authentication required)
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Protected routes (authentication required)
	protected := router.Group("/api/auth")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/logout", controllers.Logout)
		protected.GET("/me", controllers.Me)
	}
}
