package models

// MetaModel -
type MetaModel struct {
	Notes string          `bson:"notes" json:"notes"`
	Files []FileSchemaWID `bson:"files" json:"files"`
}

// NewMetaModel -
func NewMetaModel() MetaModel {
	return MetaModel{}
}
