package projects

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Remove Project
func (h *Handler) Remove(id string) (fiber.Map, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// get current document
	foundOldDoc, err := h.DB.FindByID(h.Ctx, enums.CollProjects, oid, h.Business.ID)
	if err != nil {
		return nil, err
	}
	oldDoc := foundOldDoc.(*models.ProjectModel)

	if !h.DB.DeleteByID(h.Ctx, enums.CollProjects, oldDoc.ID) {
		return nil, errors.NewHTTPError(errors.ErrDelete)
	}

	return fiber.Map{
		"project": id,
	}, nil
}
