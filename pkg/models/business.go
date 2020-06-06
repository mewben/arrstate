package models

import (
	"time"

	"github.com/mewben/realty278/internal/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BusinessModel -
type BusinessModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Domain    string             `bson:"domain" json:"domain" validate:"required,min=3,max=254"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	AreaUnits *AreaUnitsModel    `bson:"areaUnits" json:"areaUnits"`
}

// NewBusinessModel -
func NewBusinessModel() *BusinessModel {
	now := time.Now()
	return &BusinessModel{
		CreatedAt: &now,
		AreaUnits: NewAreaUnitsModel(),
	}
}

// AreaUnitsModel -
type AreaUnitsModel struct {
	Default string   `bson:"default" json:"default"`
	List    []string `bson:"list" json:"list"`
}

// NewAreaUnitsModel -
func NewAreaUnitsModel() *AreaUnitsModel {
	list := make([]string, 0)
	return &AreaUnitsModel{
		Default: enums.DefaultUnitArea,
		List:    list,
	}
}
