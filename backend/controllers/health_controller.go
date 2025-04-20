package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/habdil/notify-vital/backend/models"
	"github.com/habdil/notify-vital/backend/services"
)

// GetCurrentHealthData retrieves the latest health data for the authenticated user
func GetCurrentHealthData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	healthData, err := services.GetHealthDataForUser(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Health data not found: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": healthData})
}

// GetHealthDataHistory retrieves health data history for the authenticated user
func GetHealthDataHistory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse query parameters
	var filters models.HealthDataFilters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters: " + err.Error()})
		return
	}

	// Apply default values if not provided
	if filters.Limit <= 0 {
		filters.Limit = 30
	}

	healthDataList, err := services.GetHealthDataHistory(userID.(int), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve health data history: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": healthDataList})
}

// CreateHealthData creates a new health data entry for the authenticated user
func CreateHealthData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	var req models.HealthDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	healthData, err := services.CreateHealthData(userID.(int), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create health data: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Health data created successfully", "data": healthData})
}

// GetHealthDataSummary retrieves summary statistics for health data
func GetHealthDataSummary(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	summary, err := services.GetHealthDataSummary(userID.(int), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve health data summary: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": summary})
}

// GetHeartRateHistory retrieves heart rate history for the authenticated user
func GetHeartRateHistory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse query parameters
	var filters models.HealthDataFilters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters: " + err.Error()})
		return
	}

	// Apply default values if not provided
	if filters.Limit <= 0 {
		filters.Limit = 30
	}

	heartRateData, err := services.GetHeartRateHistory(userID.(int), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve heart rate data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": heartRateData})
}

// CreateHeartRateData creates a new heart rate data entry
func CreateHeartRateData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	var req models.HeartRateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	heartRateData, err := services.CreateHeartRateData(userID.(int), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create heart rate data: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Heart rate data created successfully", "data": heartRateData})
}

// GetStepsHistory retrieves steps history for the authenticated user
func GetStepsHistory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse query parameters
	var filters models.HealthDataFilters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters: " + err.Error()})
		return
	}

	// Apply default values if not provided
	if filters.Limit <= 0 {
		filters.Limit = 30
	}

	stepsData, err := services.GetStepsHistory(userID.(int), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve steps data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stepsData})
}

// CreateStepsData creates a new steps data entry
func CreateStepsData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	var req models.StepsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	stepsData, err := services.CreateStepsData(userID.(int), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create steps data: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Steps data created successfully", "data": stepsData})
}

// GetCaloriesHistory retrieves calories history for the authenticated user
func GetCaloriesHistory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse query parameters
	var filters models.HealthDataFilters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters: " + err.Error()})
		return
	}

	// Apply default values if not provided
	if filters.Limit <= 0 {
		filters.Limit = 30
	}

	caloriesData, err := services.GetCaloriesHistory(userID.(int), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve calories data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": caloriesData})
}

// CreateCaloriesData creates a new calories data entry
func CreateCaloriesData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	var req models.CaloriesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	caloriesData, err := services.CreateCaloriesData(userID.(int), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create calories data: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Calories data created successfully", "data": caloriesData})
}

// GetActivityStatusHistory retrieves activity status history for the authenticated user
func GetActivityStatusHistory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	// Parse query parameters
	var filters models.HealthDataFilters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters: " + err.Error()})
		return
	}

	// Apply default values if not provided
	if filters.Limit <= 0 {
		filters.Limit = 30
	}

	statusUpdates, err := services.GetActivityStatusHistory(userID.(int), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activity status data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": statusUpdates})
}

// CreateActivityStatusUpdate creates a new activity status update
func CreateActivityStatusUpdate(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	var req models.ActivityStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	statusUpdate, err := services.CreateActivityStatusUpdate(userID.(int), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activity status update: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Activity status update created successfully", "data": statusUpdate})
}
