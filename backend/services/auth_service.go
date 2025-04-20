package services

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/habdil/notify-vital/backend/config"
	"github.com/habdil/notify-vital/backend/models"
)

// HashPassword creates a bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its plaintext version
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT generates a JWT token for authentication
func GenerateJWT(userID int) (string, time.Time, error) {
	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", time.Time{}, errors.New("JWT_SECRET not set in environment")
	}

	// Get JWT expiry duration from environment
	jwtExpiry := os.Getenv("JWT_EXPIRY")
	if jwtExpiry == "" {
		jwtExpiry = "24h" // Default to 24 hours if not set
	}

	// Parse the expiry duration
	expiryDuration, err := time.ParseDuration(jwtExpiry)
	if err != nil {
		return "", time.Time{}, err
	}

	expiryTime := time.Now().Add(expiryDuration)

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expiryTime.Unix(),
		"iat":     time.Now().Unix(),
	})

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiryTime, nil
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenString string) (int, error) {
	// Get JWT secret from environment
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return 0, errors.New("JWT_SECRET not set in environment")
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return 0, err
	}

	// Validate the token and extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user ID
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("invalid user_id in token")
		}
		userID := int(userIDFloat)
		return userID, nil
	}

	return 0, errors.New("invalid token")
}

// RegisterUser registers a new user
func RegisterUser(req models.RegisterRequest) (*models.User, error) {
	// Hash the password
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create a new user in the database
	var user models.User
	err = config.DB.QueryRow(
		"INSERT INTO users (username, email, password_hash, created_at, is_active) VALUES ($1, $2, $3, $4, $5) RETURNING user_id, username, email, created_at, is_active",
		req.Username,
		req.Email,
		hashedPassword,
		time.Now(),
		true,
	).Scan(&user.UserID, &user.Username, &user.Email, &user.CreatedAt, &user.IsActive)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// LoginUser authenticates a user and returns user data and token
func LoginUser(req models.LoginRequest) (*models.User, string, time.Time, error) {
	// Find the user by email
	var user models.User
	var passwordHash string

	err := config.DB.QueryRow(
		"SELECT user_id, username, email, password_hash, created_at, is_active FROM users WHERE email = $1",
		req.Email,
	).Scan(&user.UserID, &user.Username, &user.Email, &passwordHash, &user.CreatedAt, &user.IsActive)

	if err != nil {
		return nil, "", time.Time{}, errors.New("invalid email or password")
	}

	// Check if the user is active
	if !user.IsActive {
		return nil, "", time.Time{}, errors.New("account is not active")
	}

	// Verify the password
	if !CheckPasswordHash(req.Password, passwordHash) {
		return nil, "", time.Time{}, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, expiryTime, err := GenerateJWT(user.UserID)
	if err != nil {
		return nil, "", time.Time{}, err
	}

	// Update last login time
	_, err = config.DB.Exec("UPDATE users SET last_login = $1 WHERE user_id = $2", time.Now(), user.UserID)
	if err != nil {
		// Non-critical error, just log it
		// log.Printf("Failed to update last login: %v", err)
	}

	// Create a session record
	_, err = config.DB.Exec(
		"INSERT INTO sessions (user_id, token, ip_address, issued_at, expires_at, is_valid) VALUES ($1, $2, $3, $4, $5, $6)",
		user.UserID, token, "", time.Now(), expiryTime, true,
	)
	if err != nil {
		// Non-critical error, just log it
		// log.Printf("Failed to create session: %v", err)
	}

	return &user, token, expiryTime, nil
}

// LogoutUser invalidates a user's session
func LogoutUser(token string) error {
	// Mark the session as invalid
	_, err := config.DB.Exec("UPDATE sessions SET is_valid = false WHERE token = $1", token)
	return err
}

// GetUserByID retrieves a user by ID
func GetUserByID(userID int) (*models.User, error) {
	var user models.User

	err := config.DB.QueryRow(
		"SELECT user_id, username, email, created_at, is_active FROM users WHERE user_id = $1",
		userID,
	).Scan(&user.UserID, &user.Username, &user.Email, &user.CreatedAt, &user.IsActive)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
