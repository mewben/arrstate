package models

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PersonModel -
type PersonModel struct {
	UserID         *primitive.ObjectID `bson:"userID" json:"userID"` // can be nil
	Status         string              `bson:"status" json:"status"`
	Role           string              `bson:"role" json:"role" validate:"required"`
	Email          string              `bson:"email" json:"email" validate:"email,required"`
	GivenName      string              `bson:"givenName" json:"givenName" validate:"required"`
	FamilyName     string              `bson:"familyName" json:"familyName"`
	CommissionPerc float64             `bson:"commissionPerc" json:"commissionPerc" validate:"number,min=0"`
	CustomFields   fiber.Map           `bson:"customFields" json:"customFields"`
	Address        AddressModel        `bson:"address" json:"address"`
	// Extended
	BaseModel `bson:",inline"`
	MetaModel `bson:",inline"`
}

// NewPersonModel -
func NewPersonModel(arg ...primitive.ObjectID) *PersonModel {
	return &PersonModel{
		BaseModel: NewBaseModel(arg...),
		Address:   NewAddressModel(),
		MetaModel: NewMetaModel(),
	}
}
