package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindByIDAndUpdate - returns the document model
// This is just a little helper to call on FindOneAndUpdate
func (s *Service) FindByIDAndUpdate(ctx context.Context, collectionName string, oid primitive.ObjectID, update interface{}, opts ...*options.FindOneAndUpdateOptions) (interface{}, error) {

	filter := bson.D{
		{
			Key:   "_id",
			Value: oid,
		},
	}

	return s.FindOneAndUpdate(ctx, collectionName, filter, update, opts...)
}
