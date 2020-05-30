package models

import "time"

// BusinessModel -
type BusinessModel struct {
	ID        string     `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string     `bson:"name" json:"name"`
	Domain    string     `bson:"domain" json:"domain"`
	CreatedAt *time.Time `bson:"createdAt" json:"createdAt"`
}

// NewBusinessModel -
func NewBusinessModel() *BusinessModel {
	now := time.Now()
	return &BusinessModel{
		CreatedAt: &now,
	}
}
