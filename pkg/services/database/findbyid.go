package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindByID - returns the document model
// gets from cache first then db
func (s *Service) FindByID(ctx context.Context, collectionName, id string) interface{} {
	// TODO: hit cache first
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("error findByID", err)
		return nil
	}

	filter := bson.D{
		{
			Key: "_id",
			Value: oid,
		},
	}

	return s.FindOne(ctx, collectionName, filter)
}