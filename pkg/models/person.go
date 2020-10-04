package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PersonModel -
type PersonModel struct {
	UserID         *primitive.ObjectID `bson:"userID" json:"userID"` // can be nil
	Status         string              `bson:"status" json:"status"`
	Role           []string            `bson:"role" json:"role" validate:"required"`
	Email          string              `bson:"email" json:"email" validate:"email,required"`
	Name           PersonName          `bson:"name" json:"name"`
	Avatar         *FileSchemaWID      `bson:"avatar" json:"avatar"`
	CommissionPerc int64               `bson:"commissionPerc" json:"commissionPerc" validate:"number,min=0,max=10000"`
	CustomFields   fiber.Map           `bson:"customFields" json:"customFields"`
	Address        AddressModel        `bson:"address" json:"address"`
	Locale         Locale              `bson:"locale" json:"locale"`
	// Extended
	BaseModel `bson:",inline"`
	MetaModel `bson:",inline"`
}

// PersonName -
type PersonName struct {
	First  string `bson:"first" json:"first" validate:"required"`
	Last   string `bson:"last" json:"last"`
	Middle string `bson:"middle" json:"middle"`
}

// Locale -
type Locale struct {
	Language        string `bson:"language" json:"language"`
	TimeZone        string `bson:"timeZone" json:"timeZone"`
	WeekStartDay    string `bson:"weekStartDay" json:"weekStartDay"`
	DateFormat      string `bson:"dateFormat" json:"dateFormat"`
	TimeFormat      string `bson:"timeFormat" json:"timeFormat"`
	TimestampFormat string `bson:"timestampFormat" json:"timestampFormat"`
}

// NewPersonModel -
func NewPersonModel(arg ...primitive.ObjectID) *PersonModel {
	return &PersonModel{
		BaseModel: NewBaseModel(arg...),
		Address:   NewAddressModel(),
		MetaModel: NewMetaModel(),
	}
}
