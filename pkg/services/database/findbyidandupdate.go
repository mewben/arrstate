package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindByIDAndUpdate - returns the document model
// This is just a little helper to call on FindOneAndUpdate
func (s *Service) FindByIDAndUpdate(ctx context.Context, collectionName, id string, update interface{}, opts ...*options.FindOneAndUpdateOptions) interface{} {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("error findByID", err)
		return nil
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: oid,
		},
	}

	return s.FindOneAndUpdate(ctx, collectionName, filter, update, opts...)
}
