package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find -
func (s *Service) Find(ctx context.Context, collectionName string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return s.DB.Collection(collectionName).Find(ctx, filter, opts...)
}
