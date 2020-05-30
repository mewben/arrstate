package models

// AddressModel -
type AddressModel struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	ZipCode string `bson:"zipCode" json:"zipCode"`
	State   string `bson:"state" json:"state"`
	Country string `bson:"country" json:"country"` // CountryCode
}

// NewAddressModel -
func NewAddressModel() *AddressModel {
	return &AddressModel{}
}
