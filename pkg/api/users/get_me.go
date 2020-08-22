package users

import (
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetMe -
func (h *Handler) GetMe() (*models.MeModel, error) {
	response := models.NewMeModel(h.User.ID, h.Business.ID)

	response.CurrentUser.User = h.User

	// get the person models with this userID
	// this is to get the businessIDs from this person
	filter := bson.D{
		{
			Key:   "userID",
			Value: h.User.ID,
		},
	}
	peopleCursor, err := h.DB.Find(h.Ctx, enums.CollPeople, filter)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound, err)
	}
	people := make([]*models.PersonModel, 0)
	if err = peopleCursor.All(h.Ctx, &people); err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound, err)
	}

	var businessIDs []primitive.ObjectID
	for _, person := range people {
		if person.BusinessID == h.Business.ID {
			response.CurrentUser.Person = person
		}
		businessIDs = append(businessIDs, person.BusinessID)
	}

	// get the userBusinesses
	filter = bson.D{
		{
			Key: "_id",
			Value: bson.D{
				{
					Key:   "$in",
					Value: businessIDs,
				},
			},
		},
	}
	businessesCursor, err := h.DB.Find(h.Ctx, enums.CollBusinesses, filter)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound, err)
	}
	businesses := make([]*models.BusinessModel, 0)
	if err = businessesCursor.All(h.Ctx, &businesses); err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound, err)
	}

	for _, b := range businesses {
		if b.ID == h.Business.ID {
			response.CurrentBusiness = b
		}
	}
	response.UserBusinesses = businesses

	return response, nil
}
