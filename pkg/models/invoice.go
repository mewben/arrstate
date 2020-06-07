package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InvoiceModel -
// For the payment schedules of the client
// This should be auto generated when attaching a client to a lot
type InvoiceModel struct {
	ClientLotID   *primitive.ObjectID `bson:"clientLotID" json:"clientLotID"`
	InvoiceID     string              `bson:"invoiceID" json:"invoiceID"` // some sequence or edited
	Status        string              `bson:"status" json:"status"`
	Discount      []DiscountModel     `bson:"discount" json:"discount"`
	Tax           []TaxModel          `bson:"tax" json:"tax"`
	SubTotal      float32             `bson:"subTotal" json:"subTotal" validate:"number,min=0'`
	TotalDiscount float32             `bson:"totalDiscount" json:"totalDiscount" validate:"number,min=0'`
	TotalTax      float32             `bson:"totalTax" json:"totalTax" validate:"number,min=0'`
	Total         float32             `bson:"total" json:"total" validate:"required,number,min=0'`
	IssueDate     *time.Time          `bson:"issueDate" json:"issueDate"`
	DueDate       *time.Time          `bson:"dueDate" json:"dueDate"`
	PaidAt        *time.Time          `bson:"paidAt" json:"paidAt"`
	PaidBy        string              `bson:"paidBy" json:"paidBy"`
	CancelledAt   *time.Time          `bson:"cancelledAt" json:"cancelledAt"`
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
