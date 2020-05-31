package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne - returns the whole inserted document
func (s *Service) InsertOne(ctx context.Context, collectionName string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	result, err := s.DB.Collection(collectionName).InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}

	// Auto FindOne to save to cache
	filter := bson.D{
		{
			Key:   "_id",
			Value: result.InsertedID,
		},
	}
	s.FindOne(ctx, collectionName, filter)

	return result, err
}
