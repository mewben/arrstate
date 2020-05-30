package models

// MetaModel -
type MetaModel struct {
	Notes  string       `bson:"notes" json:"notes"`
	Images []ImageModel `bson:"images" json:"images"`
}

// NewMetaModel -
func NewMetaModel() *MetaModel {
	return &MetaModel{}
}
