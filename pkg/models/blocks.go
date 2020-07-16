package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// BlockModel -
type BlockModel struct {
	Type       string             `bson:"type" json:"type"`
	EntityType string             `bson:"entityType" json:"entityType"`
	EntityID   primitive.ObjectID `bson:"entityID" json:"entityID"`
	BaseModel  `bson:",inline"`
}
