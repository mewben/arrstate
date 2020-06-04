package models

import (
	"time"

	"github.com/mewben/realty278/internal/enums"
)

// BusinessModel -
type BusinessModel struct {
	ID        string          `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string          `bson:"name" json:"name"`
	Domain    string          `bson:"domain" json:"domain"`
	CreatedAt *time.Time      `bson:"createdAt" json:"createdAt"`
	AreaUnits *AreaUnitsModel `bson:"areaUnits" json:"areaUnits"`
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
