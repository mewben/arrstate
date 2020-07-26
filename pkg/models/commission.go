package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommissionModel -
type CommissionModel struct {
	InvoiceID primitive.ObjectID `bson:"invoiceID" json:"invoiceID"`
	AgentID primitive.ObjectID `bson:"agentID" json:"agentID"`
	Status         string               `bson:"status" json:"status"`
	Total       int64      `bson:"total" json:"total" validate:"required,number,min=0"`
	IssueDate   *time.Time `bson:"issueDate" json:"issueDate"`
	PaidAt      *time.Time `bson:"paidAt" json:"paidAt"`

	// Extended
	BaseModel     `bson:",inline"`
	CurrencyModel `bson:",inline"`
	MetaModel     `bson:",inline"`
}