package auth

import (
	"context"
	"errors"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/notify-vital/backend/internal/models"
)

// FirebaseAuth provides authentication functionality using Firebase Auth
type FirebaseAuth struct {
	client *auth.Client
}

// NewFirebaseAuth creates a new Firebase authentication client
func NewFirebaseAuth(app *firebase.App) (*FirebaseAuth, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return &FirebaseAuth{
		client: client,
	}, nil
}

// CreateUser creates a new user in Firebase Authentication
func (fa *FirebaseAuth) CreateUser(ctx context.Context, req *models.SignUpRequest) (*models.User, error) {
	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password).
		DisplayName(req.DisplayName)

	firebaseUser, err := fa.client.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UID:         firebaseUser.UID,
		Email:       firebaseUser.Email,
		DisplayName: firebaseUser.DisplayName,
		Provider:    "password",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Gender:      req.Gender,
		Height:      req.Height,
		Weight:      req.Weight,
	}

	return user, nil
}

// GetUserByUID gets a user by UID from Firebase Authentication
func (fa *FirebaseAuth) GetUserByUID(ctx context.Context, uid string) (*models.User, error) {
	firebaseUser, err := fa.client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UID:         firebaseUser.UID,
		Email:       firebaseUser.Email,
		DisplayName: firebaseUser.DisplayName,
		PhotoURL:    firebaseUser.PhotoURL,
		PhoneNumber: firebaseUser.PhoneNumber,
		UpdatedAt:   time.Now(),
	}

	return user, nil
}

// UpdateUser updates a user in Firebase Authentication
func (fa *FirebaseAuth) UpdateUser(ctx context.Context, uid string, req *models.UpdateProfileRequest) error {
	params := (&auth.UserToUpdate{})

	if req.DisplayName != "" {
		params = params.DisplayName(req.DisplayName)
	}

	if req.PhotoURL != "" {
		params = params.PhotoURL(req.PhotoURL)
	}

	_, err := fa.client.UpdateUser(ctx, uid, params)
	return err
}

// DeleteUser deletes a user from Firebase Authentication
func (fa *FirebaseAuth) DeleteUser(ctx context.Context, uid string) error {
	return fa.client.DeleteUser(ctx, uid)
}

// VerifyIDToken verifies the given ID token
func (fa *FirebaseAuth) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	if idToken == "" {
		return nil, errors.New("id token is empty")
	}

	return fa.client.VerifyIDToken(ctx, idToken)
}

// CreateCustomToken creates a custom token for the given UID
func (fa *FirebaseAuth) CreateCustomToken(ctx context.Context, uid string) (string, error) {
	return fa.client.CustomToken(ctx, uid)
}
