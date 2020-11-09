package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Aggregate -
func (s *Service) Aggregate(ctx context.Context, collectionName string, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return s.DB.Collection(collectionName).Aggregate(ctx, pipeline, opts...)
}
