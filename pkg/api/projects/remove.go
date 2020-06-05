package projects

import (
	"github.com/gofiber/fiber"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Remove Project
func (h *Handler) Remove(id string) (fiber.Map, error) {
	// get current document
	foundOldDoc := h.DB.FindByID(h.Ctx, enums.CollProjects, id, h.Business.ID)
	if foundOldDoc == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	oldDoc := foundOldDoc.(*models.ProjectModel)

	if !h.DB.DeleteByID(h.Ctx, enums.CollProjects, oldDoc.ID) {
		return nil, errors.NewHTTPError(errors.ErrDelete)
	}

	return fiber.Map{
		"project": id,
	}, nil
}
