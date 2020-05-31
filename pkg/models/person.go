package models

// PersonModel -
type PersonModel struct {
	UserID     string `bson:"userID" json:"userID"`
	Status     string `bson:"status" json:"status"`
	Role       string `bson:"role" json:"role"`
	GivenName  string `bson:"givenName" json:"givenName"`
	FamilyName string `bson:"familyName" json:"familyName"`
	// Extended
	*BaseModel    `bson:",inline"`
	*AddressModel `bson:",inline"`
	*MetaModel    `bson:",inline"`
}

// NewPersonModel -
func NewPersonModel() *PersonModel {
	return &PersonModel{
		BaseModel:    NewBaseModel(),
		AddressModel: NewAddressModel(),
		MetaModel:    NewMetaModel(),
	}
}
