package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/habdil/notify-vital/backend/controllers"
	"github.com/habdil/notify-vital/backend/middleware"
)

// SetupHealthRoutes configures all health data related routes
func SetupHealthRoutes(router *gin.Engine) {
	// All health data routes require authentication
	health := router.Group("/api/health")
	health.Use(middleware.AuthMiddleware())
	{
		// Main health data endpoints
		health.GET("/current", controllers.GetCurrentHealthData)
		health.GET("/history", controllers.GetHealthDataHistory)
		health.POST("/record", controllers.CreateHealthData)
		health.GET("/summary", controllers.GetHealthDataSummary)

		// Heart rate specific endpoints
		heartRate := health.Group("/heart-rate")
		{
			heartRate.GET("/history", controllers.GetHeartRateHistory)
			heartRate.POST("/record", controllers.CreateHeartRateData)
		}

		// Steps specific endpoints
		steps := health.Group("/steps")
		{
			steps.GET("/history", controllers.GetStepsHistory)
			steps.POST("/record", controllers.CreateStepsData)
		}

		// Calories specific endpoints
		calories := health.Group("/calories")
		{
			calories.GET("/history", controllers.GetCaloriesHistory)
			calories.POST("/record", controllers.CreateCaloriesData)
		}

		// Activity status endpoints
		activity := health.Group("/activity")
		{
			activity.GET("/history", controllers.GetActivityStatusHistory)
			activity.POST("/status", controllers.CreateActivityStatusUpdate)
		}
	}
}
