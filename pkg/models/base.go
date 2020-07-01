package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BaseModel -
type BaseModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	BusinessID primitive.ObjectID `bson:"businessID" json:"businessID"`
	CreatedAt  *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt  *time.Time         `bson:"updatedAt" json:"updatedAt"`
	CreatedBy  primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	UpdatedBy  primitive.ObjectID `bson:"updatedBy" json:"updatedBy"`
}

// NewBaseModel autopopulates createdAt and updatedAt
func NewBaseModel(args ...primitive.ObjectID) BaseModel {
	now := time.Now()
	model := BaseModel{
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	for i, arg := range args {
		if !arg.IsZero() {
			if i == 0 {
				model.CreatedBy = arg
				model.UpdatedBy = arg
			} else if i == 1 {
				model.BusinessID = arg
			}
		}
	}

	return model
}
