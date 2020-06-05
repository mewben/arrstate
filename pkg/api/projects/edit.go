package projects

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Edit Project
func (h *Handler) Edit(id string, data *Payload) (*models.ProjectModel, error) {

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

	doc := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollProjects, id, op)
	if doc == nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate)
	}

	project := doc.(*models.ProjectModel)

	return project, nil

}
