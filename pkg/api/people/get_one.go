package people

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// GetOne person
// id = "current" or personID
func (h *Handler) GetOne(id string) (*models.PersonModel, error) {
	if id == "current" {
		filter := bson.D{
			{
				Key:   "userID",
				Value: h.User.ID,
			},
			{
				Key:   "businessID",
				Value: h.Business.ID,
			},
		}
		personFound, err := h.DB.FindOne(h.Ctx, enums.CollPeople, filter)
		if err != nil {
			return nil, err
		}
		if personFound == nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundPerson)
		}
		return personFound.(*models.PersonModel), nil

	} else {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, errors.NewHTTPError(errors.ErrInputInvalid)
		}

		doc, err := h.DB.FindByID(h.Ctx, enums.CollPeople, oid, h.Business.ID)
		if err != nil {
			return nil, err
		}
		return doc.(*models.PersonModel), nil
	}
}
