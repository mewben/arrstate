package properties

import (
	"log"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/utils"
)

// Create property
func (h *Handler) Create(data *Payload) (*models.PropertyModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}
	// validate type
	if !utils.Contains(allowedPropertyTypes, data.Type) {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// get project if exists
	if data.ProjectID != nil {
		foundProject := h.DB.FindByID(h.Ctx, enums.CollProjects, *data.ProjectID, h.Business.ID)
		if foundProject == nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundProject)
		}
	}

	property := models.NewPropertyModel(h.User.ID, h.Business.ID)
	property.ProjectID = data.ProjectID
	property.Name = data.Name
	property.Type = data.Type
	property.Area = data.Area
	property.Price = data.Price
	property.PriceAddon = data.PriceAddon
	property.MetaModel = data.MetaModel

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollProperties, property)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	// TODO: create hooks
	property = doc.(*models.PropertyModel)

	return property, nil
}
