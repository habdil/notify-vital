package models

import "time"

// HealthData represents the main health data for dashboard display
type HealthData struct {
	DataID             int       `json:"data_id"`
	UserID             int       `json:"user_id"`
	DeviceID           *int      `json:"device_id"`
	Timestamp          time.Time `json:"timestamp"`
	HeartRate          *int      `json:"heart_rate"`
	Steps              *int      `json:"steps"`
	CaloriesBurned     *int      `json:"calories_burned"`
	ActivityStatus     string    `json:"activity_status"`
	ActivityGaugeValue float64   `json:"activity_gauge_value"`
	CreatedAt          time.Time `json:"created_at"`
}

// HeartRateData represents heart rate history
type HeartRateData struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	DeviceID     *int      `json:"device_id"`
	Timestamp    time.Time `json:"timestamp"`
	HeartRate    int       `json:"heart_rate"`
	ActivityType *string   `json:"activity_type"`
	CreatedAt    time.Time `json:"created_at"`
}

// StepsData represents steps/distance history
type StepsData struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	DeviceID   *int      `json:"device_id"`
	Timestamp  time.Time `json:"timestamp"`
	StepsCount int       `json:"steps_count"`
	Distance   *float64  `json:"distance"`
	CreatedAt  time.Time `json:"created_at"`
}

// CaloriesData represents calories history
type CaloriesData struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	DeviceID       *int      `json:"device_id"`
	Timestamp      time.Time `json:"timestamp"`
	CaloriesBurned int       `json:"calories_burned"`
	ActivityType   *string   `json:"activity_type"`
	CreatedAt      time.Time `json:"created_at"`
}

// ActivityStatusUpdate represents status changes
type ActivityStatusUpdate struct {
	ID                 int       `json:"id"`
	UserID             int       `json:"user_id"`
	Timestamp          time.Time `json:"timestamp"`
	PreviousStatus     *string   `json:"previous_status"`
	CurrentStatus      string    `json:"current_status"`
	StatusChangeReason *string   `json:"status_change_reason"`
	CreatedAt          time.Time `json:"created_at"`
}

// HealthDataRequest is used for creating or updating health data
type HealthDataRequest struct {
	DeviceID           *int    `json:"device_id"`
	HeartRate          *int    `json:"heart_rate"`
	Steps              *int    `json:"steps"`
	CaloriesBurned     *int    `json:"calories_burned"`
	ActivityStatus     string  `json:"activity_status" binding:"required"`
	ActivityGaugeValue float64 `json:"activity_gauge_value" binding:"required"`
}

// HeartRateRequest is used for creating heart rate data
type HeartRateRequest struct {
	DeviceID     *int    `json:"device_id"`
	HeartRate    int     `json:"heart_rate" binding:"required"`
	ActivityType *string `json:"activity_type"`
}

// StepsRequest is used for creating steps data
type StepsRequest struct {
	DeviceID   *int     `json:"device_id"`
	StepsCount int      `json:"steps_count" binding:"required"`
	Distance   *float64 `json:"distance"`
}

// CaloriesRequest is used for creating calories data
type CaloriesRequest struct {
	DeviceID       *int    `json:"device_id"`
	CaloriesBurned int     `json:"calories_burned" binding:"required"`
	ActivityType   *string `json:"activity_type"`
}

// ActivityStatusRequest is used for creating activity status updates
type ActivityStatusRequest struct {
	PreviousStatus     *string `json:"previous_status"`
	CurrentStatus      string  `json:"current_status" binding:"required"`
	StatusChangeReason *string `json:"status_change_reason"`
}

// HealthDataFilters represents query parameters for filtering health data
type HealthDataFilters struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Limit     int    `form:"limit,default=30"`
	Offset    int    `form:"offset,default=0"`
}

// HealthDataSummary represents summary statistics for health data
type HealthDataSummary struct {
	AverageHeartRate     float64        `json:"average_heart_rate"`
	TotalSteps           int            `json:"total_steps"`
	TotalCaloriesBurned  int            `json:"total_calories_burned"`
	ActivityDistribution map[string]int `json:"activity_distribution"`
	PeriodStart          time.Time      `json:"period_start"`
	PeriodEnd            time.Time      `json:"period_end"`
}
