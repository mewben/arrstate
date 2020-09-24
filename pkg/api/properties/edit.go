package properties

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Edit Property
func (h *Handler) Edit(data *Payload) (*models.PropertyModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	// validate type
	if !utils.Contains(allowedPropertyTypes, data.Type) {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// get current document
	foundOldDoc, err := h.DB.FindByID(h.Ctx, enums.CollProperties, data.ID, h.Business.ID)
	if err != nil {
		return nil, err
	}
	oldDoc := foundOldDoc.(*models.PropertyModel)

	log.Println("preparing upd", oldDoc)

	// prepare fields to be $set
	upd := fiber.Map{
		"name":       data.Name,
		"type":       data.Type,
		"area":       data.Area,
		"price":      data.Price,
		"priceAddon": data.PriceAddon,
		"notes":      data.Notes,
		"updatedBy":  h.User.ID,
	}

	op := bson.D{
		{
			Key:   "$set",
			Value: upd,
		},
		{
			Key: "$currentDate",
			Value: fiber.Map{
				"updatedAt": true,
			},
		},
	}

	doc, err := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollProperties, oldDoc.ID, op)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate)
	}

	property := doc.(*models.PropertyModel)

	return property, nil
}
