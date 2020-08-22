package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Count -
func (s *Service) Count(ctx context.Context, collectionName string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return s.DB.Collection(collectionName).CountDocuments(ctx, filter, opts...)
}
