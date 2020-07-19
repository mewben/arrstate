package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BlockModel -
type BlockModel struct {
	Type       string             `bson:"type" json:"type"`
	EntityType string             `bson:"entityType" json:"entityType"`
	EntityID   primitive.ObjectID `bson:"entityID" json:"entityID"`
	BaseModel  `bson:",inline"`
}

// InvoiceItemBlockModel -
type InvoiceItemBlockModel struct {
	Title          string  `bson:"title" json:"title" validate:"required"`
	Description    string  `bson:"description" json:"description"`
	Amount         int64   `bson:"amount" json:"amount" validate:"number,min=0"`
	Quantity       float64 `bson:"quantity" json:"quantity" validate:"number,min=0"`
	Tax            int64   `bson:"tax" json:"tax" validate:"number,min=0"`
	TaxAmount      int64   `bson:"taxAmount" json:"taxAmount"`
	Discount       string  `bson:"discount" json:"discount" validate:"numberOrPercentage"`
	DiscountAmount int64   `bson:"discountAmount" json:"discountAmount"`
	Total          int64   `bson:"total" json:"total"`

	BlockModel `bson:",inline"`
}

// NewBlockModel -
func NewBlockModel(arg ...primitive.ObjectID) *BlockModel {
	return &BlockModel{
		BaseModel: NewBaseModel(arg...),
	}
}

// NewInvoiceItemBlockModel -
func NewInvoiceItemBlockModel(arg ...primitive.ObjectID) *InvoiceItemBlockModel {
	return &InvoiceItemBlockModel{
		BlockModel: *NewBlockModel(arg...),
	}
}
