package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/models"
)

// FindBusiness -
// This is a special function to get a business by _id.
// We do this for quickly caching business
func (s *Service) FindBusiness(ctx context.Context, oid primitive.ObjectID) *models.BusinessModel {
	// TODO: hit cache first
	business := models.NewBusinessModel()
	filter := bson.D{
		{
			Key:   "_id",
			Value: oid,
		},
	}
	err := s.DB.Collection(enums.CollBusinesses).FindOne(ctx, filter).Decode(&business)
	if err != nil {
		return nil
	}

	return business

}
