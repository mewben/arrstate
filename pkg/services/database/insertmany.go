package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertMany -
func (s *Service) InsertMany(ctx context.Context, collectionName string, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return s.DB.Collection(collectionName).InsertMany(ctx, documents, opts...)
}
