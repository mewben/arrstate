package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindByID - returns the document model
// gets from cache first then db
func (s *Service) FindByID(ctx context.Context, collectionName string, oid, businessID primitive.ObjectID) (interface{}, error) {
	// TODO: hit cache first

	filter := bson.D{
		{
			Key:   "_id",
			Value: oid,
		},
		{
			Key:   "businessID",
			Value: businessID,
		},
	}

	return s.FindOne(ctx, collectionName, filter)
}
