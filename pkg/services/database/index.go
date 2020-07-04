package database

// This acts as a central db manipulation service.
// This is to employ caching.
// Example. When inserting a document, it saves to cache after insert

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Service -
type Service struct {
	DB *mongo.Database
}

// NewService -
func NewService(db *mongo.Database) *Service {
	return &Service{
		DB: db,
	}
}
