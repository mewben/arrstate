package businesses

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// Get businesses for the current user
func (h *Handler) Get() (*ResponseList, error) {
	// 1. get the businessIDs from people where userID
	filter := bson.D{
		{
			Key:   "userID",
			Value: h.User.ID,
		},
	}
	opts := options.Find().SetProjection(bson.M{"businessID": 1})
	cursor, err := h.DB.Find(h.Ctx, enums.CollPeople, filter, opts)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound, err)
	}

	people := make([]*models.PersonModel, 0)
	if err = cursor.All(h.Ctx, &people); err != nil {
		log.Println("err cursor", err)
		return nil, err
	}

	var businessIDs []primitive.ObjectID
	for _, person := range people {
		businessIDs = append(businessIDs, person.BusinessID)
	}

	// 2. get the list of businesses where $in businessIDs
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
	cursor, err = h.DB.Find(h.Ctx, enums.CollBusinesses, filter)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound, err)
	}

	businesses := make([]*models.BusinessModel, 0)
	if err = cursor.All(h.Ctx, &businesses); err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound, err)
	}

	response := &ResponseList{
		Total: len(businesses),
		Data:  businesses,
	}

	return response, nil

}
