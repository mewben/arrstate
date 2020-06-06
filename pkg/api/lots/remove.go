package lots

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Remove Lot
func (h *Handler) Remove(id string) (fiber.Map, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// get current document
	foundOldDoc := h.DB.FindByID(h.Ctx, enums.CollLots, oid, h.Business.ID)
	if foundOldDoc == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	oldDoc := foundOldDoc.(*models.LotModel)

	if !h.DB.DeleteByID(h.Ctx, enums.CollLots, oldDoc.ID) {
		return nil, errors.NewHTTPError(errors.ErrDelete)
	}

	return fiber.Map{
		"lot": id,
	}, nil
}
