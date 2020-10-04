package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
)

// ProjectModel -
type ProjectModel struct {
	Name    string       `bson:"name" json:"name" validate:"required"`
	Address AddressModel `bson:"address" json:"address"`
	Area    float64      `bson:"area" json:"area" validate:"number,min=0"`
	Unit    string       `bson:"unit" json:"unit"`
	// Extended
	BaseModel     `bson:",inline"`
	CurrencyModel `bson:",inline"`
	MetaModel     `bson:",inline"`
	// Computed
	PropertyIDs []primitive.ObjectID `bson:"propertyIDs" json:"propertyIDs"`
	ClientIDs   []primitive.ObjectID `bson:"clientIDs" json:"clientIDs"`
	AgentIDs    []primitive.ObjectID `bson:"agentIDs" json:"agentIDs"`
}

// NewProjectModel -
func NewProjectModel(arg ...primitive.ObjectID) *ProjectModel {
	return &ProjectModel{
		BaseModel:     NewBaseModel(arg...),
		CurrencyModel: NewCurrencyModel(),
		MetaModel:     NewMetaModel(),
		Address:       NewAddressModel(),
		Unit:          enums.DefaultUnitArea,
	}
}
