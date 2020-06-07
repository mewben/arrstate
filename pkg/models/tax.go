package models

// TaxModel -
type TaxModel struct {
	Title string `bson:"title" json:"title"`
	// Value e.g. 12 for 12%
	Value float32 `bson:"value" json:"value" validate:"number,min=0"`
	// Amount computed from Value
	Amount float32 `bson:"amount" json:"amount" validate:"number,min=0"`
}
