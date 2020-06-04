package projects

import (
	"log"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Payload -
type Payload struct {
	Name    string               `json:"name" validate:"required"`
	Address *models.AddressModel `json:"address"`
	Area    float32              `json:"area" validate:"number,min=0"`
	Unit    string               `json:"unit"`
	*models.MetaModel
}

// Create project
func (h *Handler) Create(data *Payload) (*models.ProjectModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	// Get business
	foundBusiness := h.DB.FindByID(h.Ctx, enums.CollBusinesses, h.BusinessID)
	if foundBusiness == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFoundBusiness)
	}
	business := foundBusiness.(*models.BusinessModel)

	project := models.NewProjectModel(h.UserID, h.BusinessID)
	project.Name = data.Name
	project.Address = data.Address
	project.Area = data.Area
	project.Unit = data.Unit
	if data.Unit == "" {
		project.Unit = business.AreaUnits.Default
	}
	project.MetaModel = data.MetaModel

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollProjects, project)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	// TODO: create hooks
	project = doc.(*models.ProjectModel)

	return project, nil
}
