package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// DeleteOne - returns the document model
// caching is handled in the FindById
func (s *Service) DeleteOne(ctx context.Context, collectionName string, filter interface{}, opts ...*options.DeleteOptions) bool {
	result, err := s.DB.Collection(collectionName).DeleteOne(ctx, filter, opts...)

	if err != nil {
		log.Println("error deleteone", err)
		return false
	}

	if result.DeletedCount < 1 {
		log.Println("deletedcount", result.DeletedCount)
		return false
	}

	return true
}
