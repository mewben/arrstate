package projects

import (
	"log"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Create project
func (h *Handler) Create(data *Payload) (*models.ProjectModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	project := models.NewProjectModel(h.User.ID, h.Business.ID)
	project.Name = data.Name
	project.Address = data.Address
	project.Area = data.Area
	project.Unit = data.Unit
	if data.Unit == "" {
		project.Unit = h.Business.AreaUnits.Default
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
