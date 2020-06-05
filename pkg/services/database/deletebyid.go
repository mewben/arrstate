package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteByID - bool
// safe to omit businessID as it is expected that
// you find the document first with businessID
// remove from cache on success
func (s *Service) DeleteByID(ctx context.Context, collectionName string, oid primitive.ObjectID) bool {

	filter := bson.D{
		{
			Key:   "_id",
			Value: oid,
		},
	}

	success := s.DeleteOne(ctx, collectionName, filter)
	if success {
		// TODO: remove from cache
	}
	return success
}
