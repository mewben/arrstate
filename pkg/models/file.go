package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FileSchema -
type FileSchema struct {
	Title     string `bson:"title" json:"title"`
	Extension string `bson:"ext" json:"ext"`
	MimeType  string `bson:"mimeType" json:"mimeType"`
	Size      int    `bson:"size" json:"size"`
	URL       string `bson:"url" json:"url"`
}

// FileSchemaWID - fileSchema with ID
type FileSchemaWID struct {
	ID         *primitive.ObjectID `bson:"_id" json:"_id"`
	FileSchema `bson:",inline"`
	// Title     string             `bson:"title" json:"title"`
	// Extension string             `bson:"ext" json:"ext"`
	// MimeType  string             `bson:"mimeType" json:"mimeType"`
	// Size      int                `bson:"size" json:"size"`
	// URL       string             `bson:"url" json:"url"`
}

// FileModel -
type FileModel struct {
	FileSchema `bson:",inline"`
	Type       string `bson:"type" json:"type"`
	Link       string `bson:"link" json:"link" validate:"omitempty,url"`

	// Extended
	EntitySchema `bson:",inline"`
	BaseModel    `bson:",inline"`
}

// NewFileSchema -
func NewFileSchema() FileSchema {
	return FileSchema{}
}

// NewFileSchemaWID -
func NewFileSchemaWID() *FileSchemaWID {
	return &FileSchemaWID{
		FileSchema: NewFileSchema(),
	}
}

// NewFileModel -
func NewFileModel(arg ...primitive.ObjectID) *FileModel {
	return &FileModel{
		FileSchema:   NewFileSchema(),
		EntitySchema: NewEntitySchema(),
		BaseModel:    NewBaseModel(arg...),
	}
}
