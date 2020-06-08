package people

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Edit Person
func (h *Handler) Edit(data *Payload) (*models.PersonModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	if data.Role == enums.RoleOwner {
		return nil, errors.NewHTTPError(errors.ErrOwner)
	}

	// get current document
	foundOldDoc := h.DB.FindByID(h.Ctx, enums.CollPeople, data.ID, h.Business.ID)
	if foundOldDoc == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	oldDoc := foundOldDoc.(*models.PersonModel)

	// prepare fields to be $set
	upd := fiber.Map{
		"userID":         data.UserID,
		"email":          data.Email,
		"role":           data.Role,
		"givenName":      data.GivenName,
		"familyName":     data.FamilyName,
		"address":        data.Address,
		"notes":          data.Notes,
		"commissionPerc": data.CommissionPerc,
		"customFields":   data.CustomFields,
		"updatedBy":      h.User.ID,
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

	doc := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollPeople, oldDoc.ID, op)
	if doc == nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate)
	}

	person := doc.(*models.PersonModel)

	return person, nil

}