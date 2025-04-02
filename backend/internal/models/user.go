package models

import "time"

// User represents the user model
type User struct {
	UID         string    `json:"uid" firestore:"uid"`
	Email       string    `json:"email" firestore:"email"`
	DisplayName string    `json:"displayName" firestore:"displayName"`
	PhotoURL    string    `json:"photoURL,omitempty" firestore:"photoURL,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty" firestore:"phoneNumber,omitempty"`
	Provider    string    `json:"provider" firestore:"provider"`
	CreatedAt   time.Time `json:"createdAt" firestore:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" firestore:"updatedAt"`
	// Health profile fields
	DateOfBirth time.Time `json:"dateOfBirth,omitempty" firestore:"dateOfBirth,omitempty"`
	Gender      string    `json:"gender,omitempty" firestore:"gender,omitempty"`
	Height      float64   `json:"height,omitempty" firestore:"height,omitempty"` // in cm
	Weight      float64   `json:"weight,omitempty" firestore:"weight,omitempty"` // in kg
}

// SignUpRequest represents the request body for user registration
type SignUpRequest struct {
	Email       string  `json:"email" binding:"required,email"`
	Password    string  `json:"password" binding:"required,min=6"`
	DisplayName string  `json:"displayName" binding:"required"`
	Gender      string  `json:"gender" binding:"omitempty,oneof=male female other"`
	Height      float64 `json:"height" binding:"omitempty,gt=0"`
	Weight      float64 `json:"weight" binding:"omitempty,gt=0"`
}

// SignInRequest represents the request body for user login
type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest represents the request body for updating user profile
type UpdateProfileRequest struct {
	DisplayName string  `json:"displayName" binding:"omitempty"`
	PhotoURL    string  `json:"photoURL" binding:"omitempty"`
	Gender      string  `json:"gender" binding:"omitempty,oneof=male female other"`
	Height      float64 `json:"height" binding:"omitempty,gt=0"`
	Weight      float64 `json:"weight" binding:"omitempty,gt=0"`
}

// AuthResponse represents the response for authentication operations
type AuthResponse struct {
	Token       string `json:"token"`
	ExpiresIn   int    `json:"expiresIn"`
	DisplayName string `json:"displayName"`
	UID         string `json:"uid"`
	Email       string `json:"email"`
	PhotoURL    string `json:"photoURL,omitempty"`
}
