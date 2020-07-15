package properties

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// AcquisitionPayload -
type AcquisitionPayload struct {
	PropertyID *primitive.ObjectID `bson:"propertyID" json:"propertyID"`
	models.AcquisitionModel
}

// Acquire property
func (h *Handler) Acquire(data *AcquisitionPayload) (*models.PropertyModel, error) {
	// Validate payload
	if data.PropertyID == nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// get current document
	foundOldDoc := h.DB.FindByID(h.Ctx, enums.CollProperties, *data.PropertyID, h.Business.ID)
	if foundOldDoc == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	oldDoc := foundOldDoc.(*models.PropertyModel)

	// prepare acquisition model
	upd := fiber.Map{
		"status":      enums.StatusAcquired,
		"acquisition": data.AcquisitionModel,
	}

	// use update pipeline to utilize update $$NOW
	op := bson.A{
		bson.D{
			{
				Key:   "$set",
				Value: upd,
			},
		},
		bson.D{
			{
				Key: "$set",
				Value: fiber.Map{
					"updatedAt":               "$$NOW",
					"acquisition.acquiredAt":  "$$NOW",
					"acquisition.completedAt": "$$NOW",
				},
			},
		},
	}
	doc := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollProperties, oldDoc.ID, op)
	if doc == nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate)
	}

	// hooks,
	// create invoices

	property := doc.(*models.PropertyModel)

	return property, nil
}
