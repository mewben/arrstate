package models

// LotModel -
type LotModel struct {
	ProjectID  string  `bson:"projectID" json:"projectID"`
	Name       string  `bson:"name" json:"name"`
	Area       float32 `bson:"area" json:"area"`
	Price      float32 `bson:"price" json:"price"`
	PriceAddon float32 `bson:"priceAddon" json:"priceAddon"`
	// Extended
	*BaseModel `bson:",inline"`
	*MetaModel `bson:",inline"`
}

// NewLotModel -
func NewLotModel(arg ...string) *LotModel {
	return &LotModel{
		BaseModel: NewBaseModel(arg...),
		MetaModel: NewMetaModel(),
	}
}
