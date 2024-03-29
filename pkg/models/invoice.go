package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InvoiceModel -
// For the payment schedules of the client
// This should be auto generated when attaching a client to a property
type InvoiceModel struct {
	Name       string               `bson:"name" json:"name"`
	From       *FromToModel         `bson:"from" json:"from"`
	To         *FromToModel         `bson:"to" json:"to"`
	ProjectID  *primitive.ObjectID  `bson:"projectID" json:"projectID"`
	PropertyID *primitive.ObjectID  `bson:"propertyID" json:"propertyID"`
	No         int                  `bson:"no" json:"no"` // some sequence or edited
	ReceiptNo  string               `bson:"receiptNo" json:"receiptNo"`
	Status     string               `bson:"status" json:"status"`
	Blocks     []primitive.ObjectID `bson:"blocks" json:"blocks"`
	AddOrLess  []AddOrLessModel     `bson:"addOrLess" json:"addOrLess"`
	// Tax            int64                `bson:"tax" json:"tax" validate:"number,min=0"`
	// TaxAmount      int64                `bson:"taxAmount" json:"taxAmount"`
	// Discount       string               `bson:"discount" json:"discount"`
	// DiscountAmount int64                `bson:"discountAmount" json:"discountAmount"`
	SubTotal int64 `bson:"subTotal" json:"subTotal" validate:"number,min=0"`
	// TotalDiscount  int64              `bson:"totalDiscount" json:"totalDiscount" validate:"number,min=0"`
	// TotalTax       int64              `bson:"totalTax" json:"totalTax" validate:"number,min=0"`
	Total       int64      `bson:"total" json:"total" validate:"required,number,min=0"`
	IssueDate   *time.Time `bson:"issueDate" json:"issueDate"`
	DueDate     *time.Time `bson:"dueDate" json:"dueDate"`
	PaidAt      *time.Time `bson:"paidAt" json:"paidAt"`
	CancelledAt *time.Time `bson:"cancelledAt" json:"cancelledAt"`
	// Extended
	BaseModel     `bson:",inline"`
	CurrencyModel `bson:",inline"`
	MetaModel     `bson:",inline"`
}

// FromToModel -
type FromToModel struct {
	ID         *primitive.ObjectID `bson:"_id" json:"_id"`
	EntityType string              `bson:"entityType" json:"entityType"`
}

// NewInvoiceModel -
func NewInvoiceModel(arg ...primitive.ObjectID) *InvoiceModel {
	return &InvoiceModel{
		BaseModel: NewBaseModel(arg...),
		MetaModel: NewMetaModel(),
	}
}
