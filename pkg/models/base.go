package models

import "time"

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
func NewBaseModel() *BaseModel {
	now := time.Now()
	return &BaseModel{
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
