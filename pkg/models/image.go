package models

// ImageModel -
type ImageModel struct {
	ID          string `bson:"_id" json:"_id"`
	Src         string `bson:"src" json:"src"`
	Alt         string `bson:"alt" json:"alt"`
	Description string `bson:"description" json:"description"`
}
