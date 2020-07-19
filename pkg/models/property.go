package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AcquisitionModel -
type AcquisitionModel struct {
	ClientID      *primitive.ObjectID `bson:"clientID" json:"clientID"`
	AgentID       *primitive.ObjectID `bson:"agentID" json:"agentID"`
	PaymentScheme string              `bson:"paymentScheme" json:"paymentScheme"`
	PaymentPeriod string              `bson:"paymentPeriod" json:"paymentPeriod"` // monthly, yearly
	Terms         int                 `bson:"terms" json:"terms"`                 // 60 months
	AcquiredAt    *time.Time          `bson:"acquiredAt" json:"acquiredAt"`
	CompletedAt   *time.Time          `bson:"completedAt" json:"completedAt"`
	MetaModel     `bson:",inline"`
}

// PropertyModel -
type PropertyModel struct {
	ProjectID  *primitive.ObjectID `bson:"projectID" json:"projectID"`
	Name       string              `bson:"name" json:"name" validate:"required"`
	Type       string              `bson:"type" json:"type" validate:"required"`
	Area       float64             `bson:"area" json:"area" validate:"number,min=0"`
	Price      int64               `bson:"price" json:"price" validate:"number,min=0"`
	PriceAddon int64               `bson:"priceAddon" json:"priceAddon" validate:"number,min=0"`
	Status     string              `bson:"status" json:"status"`
	// Extended
	BaseModel `bson:",inline"`
	MetaModel `bson:",inline"`

	Acquisition AcquisitionModel `bson:"acquisition" json:"acquisition"`
}

// NewPropertyModel -
func NewPropertyModel(arg ...primitive.ObjectID) *PropertyModel {
	return &PropertyModel{
		BaseModel: NewBaseModel(arg...),
		MetaModel: NewMetaModel(),
	}
}
