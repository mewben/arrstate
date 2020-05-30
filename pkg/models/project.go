package models

import "github.com/mewben/realty278/internal/enums"

// ProjectModel -
type ProjectModel struct {
	Name    string       `bson:"name" json:"name"`
	Address AddressModel `bson:"address" json:"address"`
	Area    float32      `bson:"area" json:"area"`
	Unit    string       `bson:"unit" json:"unit"`
	// Extended
	*BaseModel
	*CurrencyModel
	*MetaModel
	// Computed
	LotIDs    []string `bson:"lotIDs" json:"lotIDs"`
	LotsArea  float32  `bson:"lotsArea" json:"lotsArea"`
	ClientIDs []string `bson:"clientIDs" json:"clientIDs"`
	AgentIDs  []string `bson:"agentIDs" json:"agentIDs"`
}

// NewProject -
func NewProject() *ProjectModel {
	return &ProjectModel{
		BaseModel:     NewBaseModel(),
		CurrencyModel: NewCurrencyModel(),
		MetaModel:     NewMetaModel(),
		Unit:          enums.DefaultUnitArea,
	}
}
