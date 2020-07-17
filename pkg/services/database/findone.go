package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOne - returns the document model
// caching is handled in the FindById
// but sets the cache after found
func (s *Service) FindOne(ctx context.Context, collectionName string, filter interface{}, opts ...*options.FindOneOptions) (interface{}, error) {
	result := s.DB.Collection(collectionName).FindOne(ctx, filter, opts...)

	if result.Err() != nil {
		// catch no documents in the further below for specific error message
		if result.Err() != mongo.ErrNoDocuments {
			return nil, result.Err()
		}
	}

	return DecodeSingle(result, collectionName)
}
