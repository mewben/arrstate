package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// PersonModel -
type PersonModel struct {
	UserID     primitive.ObjectID `bson:"userID" json:"userID"`
	Status     string             `bson:"status" json:"status"`
	Role       string             `bson:"role" json:"role"`
	GivenName  string             `bson:"givenName" json:"givenName"`
	FamilyName string             `bson:"familyName" json:"familyName"`
	// Extended
	BaseModel    `bson:",inline"`
	AddressModel `bson:",inline"`
	MetaModel    `bson:",inline"`
}

// NewPersonModel -
func NewPersonModel(arg ...primitive.ObjectID) *PersonModel {
	return &PersonModel{
		BaseModel:    NewBaseModel(arg...),
		AddressModel: NewAddressModel(),
		MetaModel:    NewMetaModel(),
	}
}
