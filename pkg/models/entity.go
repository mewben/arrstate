package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// EntitySchema -
type EntitySchema struct {
	EntityType *string             `bson:"entityType" json:"entityType"`
	EntityID   *primitive.ObjectID `bson:"entityID" json:"entityID"`
}

// NewEntitySchema -
func NewEntitySchema() EntitySchema {
	return EntitySchema{}
}
