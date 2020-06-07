package models

// DiscountModel -
type DiscountModel struct {
	Title string `bson:"title" json:"title"`
	// Value amount or percentage, TODO: createValidation
	Value string `bson:"value" json:"value"`
	// Amount computed from Value
	Amount float32 `bson:"amount" json:"amount" validate:"number,min=0"`
}
