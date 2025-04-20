package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/habdil/notify-vital/backend/config"
	"github.com/habdil/notify-vital/backend/models"
)

// GetHealthDataForUser retrieves the latest health data for a user
func GetHealthDataForUser(userID int) (*models.HealthData, error) {
	var healthData models.HealthData

	query := `
		SELECT data_id, user_id, device_id, timestamp, heart_rate, steps, calories_burned, 
		       activity_status, activity_gauge_value, created_at 
		FROM health_data 
		WHERE user_id = $1 
		ORDER BY timestamp DESC 
		LIMIT 1
	`

	var deviceID sql.NullInt32
	var heartRate sql.NullInt32
	var steps sql.NullInt32
	var caloriesBurned sql.NullInt32

	err := config.DB.QueryRow(query, userID).Scan(
		&healthData.DataID, &healthData.UserID, &deviceID, &healthData.Timestamp,
		&heartRate, &steps, &caloriesBurned, &healthData.ActivityStatus,
		&healthData.ActivityGaugeValue, &healthData.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no health data found for user")
		}
		return nil, err
	}

	// Handle null values
	if deviceID.Valid {
		deviceIDInt := int(deviceID.Int32)
		healthData.DeviceID = &deviceIDInt
	}

	if heartRate.Valid {
		heartRateInt := int(heartRate.Int32)
		healthData.HeartRate = &heartRateInt
	}

	if steps.Valid {
		stepsInt := int(steps.Int32)
		healthData.Steps = &stepsInt
	}

	if caloriesBurned.Valid {
		caloriesBurnedInt := int(caloriesBurned.Int32)
		healthData.CaloriesBurned = &caloriesBurnedInt
	}

	return &healthData, nil
}

// GetHealthDataHistory retrieves health data history for a user
func GetHealthDataHistory(userID int, filters models.HealthDataFilters) ([]models.HealthData, error) {
	var healthDataList []models.HealthData

	// Base query
	query := `
		SELECT data_id, user_id, device_id, timestamp, heart_rate, steps, calories_burned, 
		       activity_status, activity_gauge_value, created_at 
		FROM health_data 
		WHERE user_id = $1
	`

	// Apply date filters if provided
	args := []interface{}{userID}
	argCount := 2

	if filters.StartDate != "" {
		query += fmt.Sprintf(" AND timestamp >= $%d", argCount)
		args = append(args, filters.StartDate)
		argCount++
	}

	if filters.EndDate != "" {
		query += fmt.Sprintf(" AND timestamp <= $%d", argCount)
		args = append(args, filters.EndDate)
		argCount++
	}

	// Add order, limit and offset
	query += " ORDER BY timestamp DESC LIMIT $" + fmt.Sprintf("%d", argCount) +
		" OFFSET $" + fmt.Sprintf("%d", argCount+1)

	args = append(args, filters.Limit, filters.Offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var healthData models.HealthData
		var deviceID sql.NullInt32
		var heartRate sql.NullInt32
		var steps sql.NullInt32
		var caloriesBurned sql.NullInt32

		err := rows.Scan(
			&healthData.DataID, &healthData.UserID, &deviceID, &healthData.Timestamp,
			&heartRate, &steps, &caloriesBurned, &healthData.ActivityStatus,
			&healthData.ActivityGaugeValue, &healthData.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		// Handle null values
		if deviceID.Valid {
			deviceIDInt := int(deviceID.Int32)
			healthData.DeviceID = &deviceIDInt
		}

		if heartRate.Valid {
			heartRateInt := int(heartRate.Int32)
			healthData.HeartRate = &heartRateInt
		}

		if steps.Valid {
			stepsInt := int(steps.Int32)
			healthData.Steps = &stepsInt
		}

		if caloriesBurned.Valid {
			caloriesBurnedInt := int(caloriesBurned.Int32)
			healthData.CaloriesBurned = &caloriesBurnedInt
		}

		healthDataList = append(healthDataList, healthData)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return healthDataList, nil
}

// CreateHealthData creates a new health data entry
func CreateHealthData(userID int, data models.HealthDataRequest) (*models.HealthData, error) {
	query := `
		INSERT INTO health_data (
			user_id, device_id, timestamp, heart_rate, steps, 
			calories_burned, activity_status, activity_gauge_value
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING data_id, timestamp, created_at
	`

	var healthData models.HealthData
	healthData.UserID = userID
	healthData.DeviceID = data.DeviceID
	healthData.HeartRate = data.HeartRate
	healthData.Steps = data.Steps
	healthData.CaloriesBurned = data.CaloriesBurned
	healthData.ActivityStatus = data.ActivityStatus
	healthData.ActivityGaugeValue = data.ActivityGaugeValue

	now := time.Now()

	err := config.DB.QueryRow(
		query,
		userID,
		data.DeviceID,
		now,
		data.HeartRate,
		data.Steps,
		data.CaloriesBurned,
		data.ActivityStatus,
		data.ActivityGaugeValue,
	).Scan(&healthData.DataID, &healthData.Timestamp, &healthData.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &healthData, nil
}

// GetHealthDataSummary retrieves summary statistics for a time period
func GetHealthDataSummary(userID int, startDate, endDate string) (*models.HealthDataSummary, error) {
	query := `
		SELECT 
			COALESCE(AVG(heart_rate), 0) as avg_heart_rate,
			COALESCE(SUM(steps), 0) as total_steps,
			COALESCE(SUM(calories_burned), 0) as total_calories_burned
		FROM health_data
		WHERE user_id = $1
	`

	args := []interface{}{userID}
	argCount := 2

	if startDate != "" {
		query += fmt.Sprintf(" AND timestamp >= $%d", argCount)
		args = append(args, startDate)
		argCount++
	}

	if endDate != "" {
		query += fmt.Sprintf(" AND timestamp <= $%d", argCount)
		args = append(args, endDate)
	}

	var summary models.HealthDataSummary

	err := config.DB.QueryRow(query, args...).Scan(
		&summary.AverageHeartRate,
		&summary.TotalSteps,
		&summary.TotalCaloriesBurned,
	)

	if err != nil {
		return nil, err
	}

	// Get activity distribution
	activityQuery := `
		SELECT 
			activity_status, 
			COUNT(*) as count
		FROM health_data
		WHERE user_id = $1
	`

	if startDate != "" {
		activityQuery += " AND timestamp >= $2"
		if endDate != "" {
			activityQuery += " AND timestamp <= $3"
		}
	} else if endDate != "" {
		activityQuery += " AND timestamp <= $2"
	}

	activityQuery += " GROUP BY activity_status"

	activityRows, err := config.DB.Query(activityQuery, args...)
	if err != nil {
		return nil, err
	}
	defer activityRows.Close()

	summary.ActivityDistribution = make(map[string]int)

	for activityRows.Next() {
		var status string
		var count int

		if err := activityRows.Scan(&status, &count); err != nil {
			return nil, err
		}

		summary.ActivityDistribution[status] = count
	}

	if err = activityRows.Err(); err != nil {
		return nil, err
	}

	// Parse dates for period information
	if startDate != "" {
		parsedStart, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			summary.PeriodStart = parsedStart
		}
	}

	if endDate != "" {
		parsedEnd, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			summary.PeriodEnd = parsedEnd
		}
	}

	return &summary, nil
}

// GetHeartRateHistory retrieves heart rate history for a user
func GetHeartRateHistory(userID int, filters models.HealthDataFilters) ([]models.HeartRateData, error) {
	var heartRateDataList []models.HeartRateData

	query := `
		SELECT id, user_id, device_id, timestamp, heart_rate, activity_type, created_at
		FROM heart_rate_data
		WHERE user_id = $1
	`

	args := []interface{}{userID}
	argCount := 2

	if filters.StartDate != "" {
		query += fmt.Sprintf(" AND timestamp >= $%d", argCount)
		args = append(args, filters.StartDate)
		argCount++
	}

	if filters.EndDate != "" {
		query += fmt.Sprintf(" AND timestamp <= $%d", argCount)
		args = append(args, filters.EndDate)
		argCount++
	}

	query += " ORDER BY timestamp DESC LIMIT $" + fmt.Sprintf("%d", argCount) +
		" OFFSET $" + fmt.Sprintf("%d", argCount+1)

	args = append(args, filters.Limit, filters.Offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var heartRateData models.HeartRateData
		var deviceID sql.NullInt32
		var activityType sql.NullString

		err := rows.Scan(
			&heartRateData.ID, &heartRateData.UserID, &deviceID, &heartRateData.Timestamp,
			&heartRateData.HeartRate, &activityType, &heartRateData.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		if deviceID.Valid {
			deviceIDInt := int(deviceID.Int32)
			heartRateData.DeviceID = &deviceIDInt
		}

		if activityType.Valid {
			activityTypeStr := activityType.String
			heartRateData.ActivityType = &activityTypeStr
		}

		heartRateDataList = append(heartRateDataList, heartRateData)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return heartRateDataList, nil
}

// CreateHeartRateData adds a new heart rate data entry
func CreateHeartRateData(userID int, data models.HeartRateRequest) (*models.HeartRateData, error) {
	query := `
		INSERT INTO heart_rate_data (
			user_id, device_id, timestamp, heart_rate, activity_type
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, timestamp, created_at
	`

	var heartRateData models.HeartRateData
	heartRateData.UserID = userID
	heartRateData.DeviceID = data.DeviceID
	heartRateData.HeartRate = data.HeartRate
	heartRateData.ActivityType = data.ActivityType

	now := time.Now()

	err := config.DB.QueryRow(
		query,
		userID,
		data.DeviceID,
		now,
		data.HeartRate,
		data.ActivityType,
	).Scan(&heartRateData.ID, &heartRateData.Timestamp, &heartRateData.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &heartRateData, nil
}

// Similar functions for steps data, calories data, and activity status updates
// would follow the same pattern...

// GetStepsHistory retrieves steps history for a user
func GetStepsHistory(userID int, filters models.HealthDataFilters) ([]models.StepsData, error) {
	var stepsDataList []models.StepsData

	query := `
		SELECT id, user_id, device_id, timestamp, steps_count, distance, created_at
		FROM steps_data
		WHERE user_id = $1
	`

	args := []interface{}{userID}
	argCount := 2

	if filters.StartDate != "" {
		query += fmt.Sprintf(" AND timestamp >= $%d", argCount)
		args = append(args, filters.StartDate)
		argCount++
	}

	if filters.EndDate != "" {
		query += fmt.Sprintf(" AND timestamp <= $%d", argCount)
		args = append(args, filters.EndDate)
		argCount++
	}

	query += " ORDER BY timestamp DESC LIMIT $" + fmt.Sprintf("%d", argCount) +
		" OFFSET $" + fmt.Sprintf("%d", argCount+1)

	args = append(args, filters.Limit, filters.Offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var stepsData models.StepsData
		var deviceID sql.NullInt32
		var distance sql.NullFloat64

		err := rows.Scan(
			&stepsData.ID, &stepsData.UserID, &deviceID, &stepsData.Timestamp,
			&stepsData.StepsCount, &distance, &stepsData.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		if deviceID.Valid {
			deviceIDInt := int(deviceID.Int32)
			stepsData.DeviceID = &deviceIDInt
		}

		if distance.Valid {
			distanceFloat := distance.Float64
			stepsData.Distance = &distanceFloat
		}

		stepsDataList = append(stepsDataList, stepsData)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stepsDataList, nil
}

// CreateStepsData adds a new steps data entry
func CreateStepsData(userID int, data models.StepsRequest) (*models.StepsData, error) {
	query := `
		INSERT INTO steps_data (
			user_id, device_id, timestamp, steps_count, distance
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, timestamp, created_at
	`

	var stepsData models.StepsData
	stepsData.UserID = userID
	stepsData.DeviceID = data.DeviceID
	stepsData.StepsCount = data.StepsCount
	stepsData.Distance = data.Distance

	now := time.Now()

	err := config.DB.QueryRow(
		query,
		userID,
		data.DeviceID,
		now,
		data.StepsCount,
		data.Distance,
	).Scan(&stepsData.ID, &stepsData.Timestamp, &stepsData.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &stepsData, nil
}

// GetCaloriesHistory retrieves calories history for a user
func GetCaloriesHistory(userID int, filters models.HealthDataFilters) ([]models.CaloriesData, error) {
	var caloriesDataList []models.CaloriesData

	query := `
		SELECT id, user_id, device_id, timestamp, calories_burned, activity_type, created_at
		FROM calories_data
		WHERE user_id = $1
	`

	args := []interface{}{userID}
	argCount := 2

	if filters.StartDate != "" {
		query += fmt.Sprintf(" AND timestamp >= $%d", argCount)
		args = append(args, filters.StartDate)
		argCount++
	}

	if filters.EndDate != "" {
		query += fmt.Sprintf(" AND timestamp <= $%d", argCount)
		args = append(args, filters.EndDate)
		argCount++
	}

	query += " ORDER BY timestamp DESC LIMIT $" + fmt.Sprintf("%d", argCount) +
		" OFFSET $" + fmt.Sprintf("%d", argCount+1)

	args = append(args, filters.Limit, filters.Offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var caloriesData models.CaloriesData
		var deviceID sql.NullInt32
		var activityType sql.NullString

		err := rows.Scan(
			&caloriesData.ID, &caloriesData.UserID, &deviceID, &caloriesData.Timestamp,
			&caloriesData.CaloriesBurned, &activityType, &caloriesData.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		if deviceID.Valid {
			deviceIDInt := int(deviceID.Int32)
			caloriesData.DeviceID = &deviceIDInt
		}

		if activityType.Valid {
			activityTypeStr := activityType.String
			caloriesData.ActivityType = &activityTypeStr
		}

		caloriesDataList = append(caloriesDataList, caloriesData)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return caloriesDataList, nil
}

// CreateCaloriesData adds a new calories data entry
func CreateCaloriesData(userID int, data models.CaloriesRequest) (*models.CaloriesData, error) {
	query := `
		INSERT INTO calories_data (
			user_id, device_id, timestamp, calories_burned, activity_type
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, timestamp, created_at
	`

	var caloriesData models.CaloriesData
	caloriesData.UserID = userID
	caloriesData.DeviceID = data.DeviceID
	caloriesData.CaloriesBurned = data.CaloriesBurned
	caloriesData.ActivityType = data.ActivityType

	now := time.Now()

	err := config.DB.QueryRow(
		query,
		userID,
		data.DeviceID,
		now,
		data.CaloriesBurned,
		data.ActivityType,
	).Scan(&caloriesData.ID, &caloriesData.Timestamp, &caloriesData.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &caloriesData, nil
}

// GetActivityStatusHistory retrieves activity status history for a user
func GetActivityStatusHistory(userID int, filters models.HealthDataFilters) ([]models.ActivityStatusUpdate, error) {
	var statusUpdatesList []models.ActivityStatusUpdate

	query := `
		SELECT id, user_id, timestamp, previous_status, current_status, status_change_reason, created_at
		FROM activity_status_updates
		WHERE user_id = $1
	`

	args := []interface{}{userID}
	argCount := 2

	if filters.StartDate != "" {
		query += fmt.Sprintf(" AND timestamp >= $%d", argCount)
		args = append(args, filters.StartDate)
		argCount++
	}

	if filters.EndDate != "" {
		query += fmt.Sprintf(" AND timestamp <= $%d", argCount)
		args = append(args, filters.EndDate)
		argCount++
	}

	query += " ORDER BY timestamp DESC LIMIT $" + fmt.Sprintf("%d", argCount) +
		" OFFSET $" + fmt.Sprintf("%d", argCount+1)

	args = append(args, filters.Limit, filters.Offset)

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var statusUpdate models.ActivityStatusUpdate
		var previousStatus sql.NullString
		var statusChangeReason sql.NullString

		err := rows.Scan(
			&statusUpdate.ID, &statusUpdate.UserID, &statusUpdate.Timestamp,
			&previousStatus, &statusUpdate.CurrentStatus, &statusChangeReason,
			&statusUpdate.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		if previousStatus.Valid {
			prevStatusStr := previousStatus.String
			statusUpdate.PreviousStatus = &prevStatusStr
		}

		if statusChangeReason.Valid {
			reasonStr := statusChangeReason.String
			statusUpdate.StatusChangeReason = &reasonStr
		}

		statusUpdatesList = append(statusUpdatesList, statusUpdate)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return statusUpdatesList, nil
}

// CreateActivityStatusUpdate adds a new activity status update
func CreateActivityStatusUpdate(userID int, data models.ActivityStatusRequest) (*models.ActivityStatusUpdate, error) {
	query := `
		INSERT INTO activity_status_updates (
			user_id, timestamp, previous_status, current_status, status_change_reason
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, timestamp, created_at
	`

	var statusUpdate models.ActivityStatusUpdate
	statusUpdate.UserID = userID
	statusUpdate.PreviousStatus = data.PreviousStatus
	statusUpdate.CurrentStatus = data.CurrentStatus
	statusUpdate.StatusChangeReason = data.StatusChangeReason

	now := time.Now()

	err := config.DB.QueryRow(
		query,
		userID,
		now,
		data.PreviousStatus,
		data.CurrentStatus,
		data.StatusChangeReason,
	).Scan(&statusUpdate.ID, &statusUpdate.Timestamp, &statusUpdate.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &statusUpdate, nil
}
