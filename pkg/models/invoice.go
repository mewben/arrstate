package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InvoiceModel -
// For the payment schedules of the client
// This should be auto generated when attaching a client to a property
type InvoiceModel struct {
	From           *FromToModel         `bson:"from" json:"from"`
	To             *FromToModel         `bson:"to" json:"to"`
	ProjectID      *primitive.ObjectID  `bson:"projectID" json:"projectID"`
	PropertyID     *primitive.ObjectID  `bson:"propertyID" json:"propertyID"`
	No             string               `bson:"no" json:"no"` // some sequence or edited
	Status         string               `bson:"status" json:"status"`
	Blocks         []primitive.ObjectID `bson:"blocks" json:"blocks"`
	Tax            float64              `bson:"tax" json:"tax" validate:"number,min=0"`
	TaxAmount      float64              `bson:"taxAmount" json:"taxAmount"`
	Discount       string               `bson:"discount" json:"discount"`
	DiscountAmount float64              `bson:"discountAmount" json:"discountAmount"`
	SubTotal       float64              `bson:"subTotal" json:"subTotal" validate:"number,min=0"`
	// TotalDiscount  float64              `bson:"totalDiscount" json:"totalDiscount" validate:"number,min=0"`
	// TotalTax       float64              `bson:"totalTax" json:"totalTax" validate:"number,min=0"`
	Total       float64    `bson:"total" json:"total" validate:"required,number,min=0"`
	IssueDate   *time.Time `bson:"issueDate" json:"issueDate"`
	DueDate     *time.Time `bson:"dueDate" json:"dueDate"`
	PaidAt      *time.Time `bson:"paidAt" json:"paidAt"`
	PaidBy      string     `bson:"paidBy" json:"paidBy"`
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
