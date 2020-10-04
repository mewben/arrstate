package properties

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// Remove Property
func (h *Handler) Remove(id string) (fiber.Map, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// get current document
	foundOldDoc, err := h.DB.FindByID(h.Ctx, enums.CollProperties, oid, h.Business.ID)
	if err != nil {
		return nil, err
	}
	oldDoc := foundOldDoc.(*models.PropertyModel)

	if !h.DB.DeleteByID(h.Ctx, enums.CollProperties, oldDoc.ID) {
		return nil, errors.NewHTTPError(errors.ErrDelete)
	}

	return fiber.Map{
		"property": id,
	}, nil
}
