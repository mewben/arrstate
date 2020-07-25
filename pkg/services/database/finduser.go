package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/models"
)

// FindUser -
// This is a special function to get a user by _id.
// We do this for quickly caching user
func (s *Service) FindUser(ctx context.Context, oid primitive.ObjectID) *models.UserModel {
	// TODO: hit cache first
	user := models.NewUserModel()
	filter := bson.D{
		{
			Key:   "_id",
			Value: oid,
		},
	}
	err := s.DB.Collection(enums.CollUsers).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil
	}

	return user

}
