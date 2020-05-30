package models

// LotModel -
type LotModel struct {
	ProjectID  string  `bson:"projectID" json:"projectID"`
	Name       string  `bson:"name" json:"name"`
	Area       float32 `bson:"area" json:"area"`
	Price      float32 `bson:"price" json:"price"`
	PriceAddon float32 `bson:"priceAddon" json:"priceAddon"`
	// Extended
	*BaseModel
	*MetaModel
}

// NewLotModel -
func NewLotModel(projectID string) *LotModel {
	return &LotModel{
		ProjectID: projectID,
		BaseModel: NewBaseModel(),
		MetaModel: NewMetaModel(),
	}
}
