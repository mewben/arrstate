package people

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// EditLocale Person
func (h *Handler) EditLocale(id string, data *models.Locale) (*models.PersonModel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	if oid != h.Person.ID {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}

	// get current document
	_, err = h.DB.FindByID(h.Ctx, enums.CollPeople, oid, h.Business.ID)
	if err != nil {
		return nil, err
	}

	// prepare fields to be $set
	upd := fiber.Map{
		"locale": data,
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

	doc, err := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollPeople, oid, op)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate)
	}

	person := doc.(*models.PersonModel)

	return person, nil

}
