package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOneAndUpdate - returns the document model
func (s *Service) FindOneAndUpdate(ctx context.Context, collectionName string, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) interface{} {
	if len(opts) == 0 {
		// set to return the updated document
		opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
		opts = append(opts, opt)
	}
	result := s.DB.Collection(collectionName).FindOneAndUpdate(ctx, filter, update, opts...)

	if result.Err() != nil {
		log.Println("findoneandupdate err", result.Err())
		return nil
	}

	return DecodeSingle(result, collectionName)
}
