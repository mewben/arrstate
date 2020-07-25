package projects

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// Edit Project
func (h *Handler) Edit(data *Payload) (*models.ProjectModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	// get current document
	foundOldDoc, err := h.DB.FindByID(h.Ctx, enums.CollProjects, data.ID, h.Business.ID)
	if err != nil {
		return nil, err
	}
	oldDoc := foundOldDoc.(*models.ProjectModel)

	// prepare fields to be $set
	upd := fiber.Map{
		"name":      data.Name,
		"address":   data.Address,
		"area":      data.Area,
		"unit":      data.Unit,
		"notes":     data.Notes,
		"updatedBy": h.User.ID,
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

	doc, err := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollProjects, oldDoc.ID, op)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate, err)
	}

	project := doc.(*models.ProjectModel)

	return project, nil

}
