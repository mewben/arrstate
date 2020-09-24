package people

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/utils"
)

// Edit Person
func (h *Handler) Edit(data *Payload) (*models.PersonModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	// get current document
	foundOldDoc, err := h.DB.FindByID(h.Ctx, enums.CollPeople, data.ID, h.Business.ID)
	if err != nil {
		return nil, err
	}
	oldDoc := foundOldDoc.(*models.PersonModel)

	if utils.Contains(data.Role, enums.RoleOwner) && !utils.Contains(oldDoc.Role, enums.RoleOwner) {
		return nil, errors.NewHTTPError(errors.ErrOwner)
	}

	// prepare fields to be $set
	upd := fiber.Map{
		"userID":         data.UserID,
		"email":          data.Email,
		"role":           data.Role,
		"name":           data.Name,
		"address":        data.Address,
		"avatar":         data.Avatar,
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

	doc, err := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollPeople, oldDoc.ID, op)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate)
	}

	person := doc.(*models.PersonModel)

	return person, nil

}
