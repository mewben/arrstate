package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InvoiceModel -
// For the payment schedules of the client
// This should be auto generated when attaching a client to a property
type InvoiceModel struct {
	PropertyID    primitive.ObjectID `bson:"propertyID" json:"propertyID"`
	InvoiceName   string             `bson:"invoiceName" json:"invoiceName"` // some sequence or edited
	Status        string             `bson:"status" json:"status"`
	Blocks        []string           `bson:"blocks" json:"blocks"`
	Discount      []DiscountModel    `bson:"discount" json:"discount"`
	Tax           []TaxModel         `bson:"tax" json:"tax"`
	SubTotal      float64            `bson:"subTotal" json:"subTotal" validate:"number,min=0"`
	TotalDiscount float64            `bson:"totalDiscount" json:"totalDiscount" validate:"number,min=0"`
	TotalTax      float64            `bson:"totalTax" json:"totalTax" validate:"number,min=0"`
	Total         float64            `bson:"total" json:"total" validate:"required,number,min=0"`
	IssueDate     *time.Time         `bson:"issueDate" json:"issueDate"`
	DueDate       *time.Time         `bson:"dueDate" json:"dueDate"`
	PaidAt        *time.Time         `bson:"paidAt" json:"paidAt"`
	PaidBy        string             `bson:"paidBy" json:"paidBy"`
	CancelledAt   *time.Time         `bson:"cancelledAt" json:"cancelledAt"`
	// Extended
	BaseModel     `bson:",inline"`
	CurrencyModel `bson:",inline"`
	MetaModel     `bson:",inline"`
}

// NewInvoiceModel -
func NewInvoiceModel(arg ...primitive.ObjectID) *InvoiceModel {
	return &InvoiceModel{
		BaseModel: NewBaseModel(arg...),
		MetaModel: NewMetaModel(),
	}
}
