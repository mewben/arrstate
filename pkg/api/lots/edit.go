package lots

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Edit Lot
func (h *Handler) Edit(data *Payload) (*models.LotModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	log.Println("finding current doc", data.ID)

	// get current document
	foundOldDoc := h.DB.FindByID(h.Ctx, enums.CollLots, data.ID, h.Business.ID)
	if foundOldDoc == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	oldDoc := foundOldDoc.(*models.LotModel)

	log.Println("preparing upd", oldDoc)

	// prepare fields to be $set
	upd := fiber.Map{
		"name":       data.Name,
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

	doc := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollLots, oldDoc.ID, op)
	if doc == nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate)
	}

	lot := doc.(*models.LotModel)

	return lot, nil
}
