package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ClientPropertyModel -
// This holds the data when a client owns a property
type ClientPropertyModel struct {
	PropertyID  primitive.ObjectID  `bson:"propertyID" json:"propertyID"`
	ClientID    primitive.ObjectID  `bson:"clientID" json:"clientID"`
	AgentID     *primitive.ObjectID `bson:"agentID" json:"agentID"`
	Status      string              `bson:"status" json:"status"`
	Price       float64             `bson:"price" json:"price" validate:"number,min=0"`
	DownPayment float64             `bson:"downPayment" json:"downPayment" validate:"required,number,min=0"`
	Months      int                 `bson:"months" json:"months" validate:"required,number,min=0"`
	Monthly     float64             `bson:"monthly" json:"monthly" validate:"required,number,min=0"`
	Date        time.Time           `bson:"date" json:"date" validate:"required"`
	ApprovedBy  *primitive.ObjectID `bson:"approvedBy" json:"approvedBy"`
	// Extended
	BaseModel `bson:",inline"`
	MetaModel `bson:",inline"`
}

// NewClientPropertyModel -
func NewClientPropertyModel(arg ...primitive.ObjectID) *ClientPropertyModel {
	return &ClientPropertyModel{
		BaseModel: NewBaseModel(arg...),
		MetaModel: NewMetaModel(),
	}
}
