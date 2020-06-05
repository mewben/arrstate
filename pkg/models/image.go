package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ImageModel -
type ImageModel struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Src         string             `bson:"src" json:"src"`
	Alt         string             `bson:"alt" json:"alt"`
	Description string             `bson:"description" json:"description"`
}
