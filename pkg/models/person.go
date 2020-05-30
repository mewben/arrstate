package models

// PersonModel -
type PersonModel struct {
	UserID     string `bson:"userID" json:"userID"`
	Status     string `bson:"status" json:"status"`
	Role       string `bson:"role" json:"role"`
	GivenName  string `bson:"givenName" json:"givenName"`
	FamilyName string `bson:"familyName" json:"familyName"`
	// Extended
	*BaseModel
	*AddressModel
	*MetaModel
}

// NewPersonModel -
func NewPersonModel(userID string) *PersonModel {
	return &PersonModel{
		UserID:       userID,
		BaseModel:    NewBaseModel(),
		AddressModel: NewAddressModel(),
		MetaModel:    NewMetaModel(),
	}
}
