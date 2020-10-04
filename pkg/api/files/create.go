package files

import (
	"log"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// Create file
func (h *Handler) Create(data *Payload) (*models.FileModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	file := models.NewFileModel(h.User.ID, h.Business.ID)
	// manually assign values to screen out unwanted inserts
	file.Title = data.Title
	file.Extension = data.Extension
	file.MimeType = data.MimeType
	file.Size = data.Size
	file.URL = data.URL
	file.EntityType = data.EntityType
	file.EntityID = data.EntityID
	file.Type = data.Type
	file.Link = data.Link

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollFiles, file)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	// TODO: create hooks
	file = doc.(*models.FileModel)

	return file, nil
}
