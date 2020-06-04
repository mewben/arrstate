package models

import (
	"time"
)

// BaseModel -
type BaseModel struct {
	ID         string     `bson:"_id,omitempty" json:"_id,omitempty"`
	BusinessID string     `bson:"businessID" json:"businessID"`
	CreatedAt  *time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt  *time.Time `bson:"updatedAt" json:"updatedAt"`
	CreatedBy  string     `bson:"createdBy" json:"createdBy"`
	UpdatedBy  string     `bson:"updatedBy" json:"updatedBy"`
}

// NewBaseModel autopopulates createdAt and updatedAt
func NewBaseModel(args ...string) *BaseModel {
	now := time.Now()
	model := &BaseModel{
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	for i, arg := range args {
		if arg != "" {
			if i == 0 {
				model.CreatedBy = arg
			} else if i == 1 {
				model.BusinessID = arg
			}
		}
	}

	return model
}
