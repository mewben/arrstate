package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOne - returns the document model
// caching is handled in the FindById
// but sets the cache after found
func (s *Service) FindOne(ctx context.Context, collectionName string, filter interface{}, opts ...*options.FindOneOptions) interface{} {
	result := s.DB.Collection(collectionName).FindOne(ctx, filter, opts...)

	if result.Err() != nil {
		return nil
	}

	return DecodeSingle(result, collectionName)
}
