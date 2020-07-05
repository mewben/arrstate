package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// PropertyModel -
type PropertyModel struct {
	ProjectID  *primitive.ObjectID `bson:"projectID" json:"projectID"`
	Name       string              `bson:"name" json:"name" validate:"required"`
	Type       string              `bson:"type" json:"type" validate:"required"`
	Area       float64             `bson:"area" json:"area" validate:"number,min=0"`
	Price      float64             `bson:"price" json:"price" validate:"number,min=0"`
	PriceAddon float64             `bson:"priceAddon" json:"priceAddon" validate:"number,min=0"`
	Status     string              `bson:"status" json:"status"`
	// Extended
	BaseModel `bson:",inline"`
	MetaModel `bson:",inline"`
	// Set on Hooks
	ClientPropertyID *primitive.ObjectID `bson:"clientPropertyID" json:"clientPropertyID"`
}

// NewPropertyModel -
func NewPropertyModel(arg ...primitive.ObjectID) *PropertyModel {
	return &PropertyModel{
		BaseModel: NewBaseModel(arg...),
		MetaModel: NewMetaModel(),
	}
}
