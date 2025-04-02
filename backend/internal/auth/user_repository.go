package auth

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/notify-vital/backend/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserRepository provides methods to interact with the user collection in Firestore
type UserRepository struct {
	client     *firestore.Client
	collection string
}

// NewUserRepository creates a new user repository
func NewUserRepository(client *firestore.Client) *UserRepository {
	return &UserRepository{
		client:     client,
		collection: "users",
	}
}

// StoreUser stores a user in Firestore
func (ur *UserRepository) StoreUser(ctx context.Context, user *models.User) error {
	_, err := ur.client.Collection(ur.collection).Doc(user.UID).Set(ctx, user)
	return err
}

// GetUserByUID gets a user by UID from Firestore
func (ur *UserRepository) GetUserByUID(ctx context.Context, uid string) (*models.User, error) {
	doc, err := ur.client.Collection(ur.collection).Doc(uid).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, err
	}

	var user models.User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user in Firestore
func (ur *UserRepository) UpdateUser(ctx context.Context, uid string, req *models.UpdateProfileRequest) error {
	updates := make(map[string]interface{})
	updates["updatedAt"] = time.Now()

	if req.DisplayName != "" {
		updates["displayName"] = req.DisplayName
	}

	if req.PhotoURL != "" {
		updates["photoURL"] = req.PhotoURL
	}

	if req.Gender != "" {
		updates["gender"] = req.Gender
	}

	if req.Height > 0 {
		updates["height"] = req.Height
	}

	if req.Weight > 0 {
		updates["weight"] = req.Weight
	}

	_, err := ur.client.Collection(ur.collection).Doc(uid).Update(ctx, []firestore.Update{
		{Path: "updatedAt", Value: time.Now()},
		{Path: "displayName", Value: req.DisplayName},
		{Path: "photoURL", Value: req.PhotoURL},
		{Path: "gender", Value: req.Gender},
		{Path: "height", Value: req.Height},
		{Path: "weight", Value: req.Weight},
	})

	return err
}

// DeleteUser deletes a user from Firestore
func (ur *UserRepository) DeleteUser(ctx context.Context, uid string) error {
	_, err := ur.client.Collection(ur.collection).Doc(uid).Delete(ctx)
	return err
}
