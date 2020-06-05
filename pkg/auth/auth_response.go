package auth

import (
	"log"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthResponse response after signup/signin
func (h *Handler) AuthResponse(userID, businessID primitive.ObjectID) (*models.AuthSuccessResponse, error) {
	response := models.NewAuthSuccessResponse(userID, businessID)

	// generate jwt
	token, err := response.CurrentUser.User.GenerateJWT(userID, businessID)
	if err != nil {
		log.Println("error generate jwt")
		return nil, err
	}
	response.Token = token

	// get user
	filter := bson.D{
		{
			Key:   "_id",
			Value: userID,
		},
	}
	user := h.DB.FindOne(h.Ctx, enums.CollUsers, filter)
	response.CurrentUser.User = user.(*models.UserModel)

	// get people
	filter = bson.D{
		{
			Key:   "userID",
			Value: userID,
		},
		{
			Key:   "businessID",
			Value: businessID,
		},
	}
	peopleCursor, err := h.DB.Find(h.Ctx, enums.CollPeople, filter)
	if err != nil {
		log.Println("err find people")
		return nil, err
	}
	people := make([]*models.PersonModel, 0)
	if err = peopleCursor.All(h.Ctx, &people); err != nil {
		log.Println("err cursor people")
		return nil, err
	}

	var businessIDs []primitive.ObjectID
	for _, person := range people {
		if person.BusinessID == businessID {
			response.CurrentUser.Person = person
		}
		businessIDs = append(businessIDs, person.BusinessID)
	}

	// get business
	filter = bson.D{
		{
			Key:   "_id",
			Value: businessID,
		},
	}
	business := h.DB.FindOne(h.Ctx, enums.CollBusinesses, filter)
	response.CurrentBusiness = business.(*models.BusinessModel)

	// get user businesses
	log.Println("businessIDs", businessIDs)
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
		log.Println("err find businesses")
		return nil, err
	}
	businesses := make([]*models.BusinessModel, 0)
	if err = businessesCursor.All(h.Ctx, &businesses); err != nil {
		log.Println("err cursor businesses")
		return nil, err
	}
	response.UserBusinesses = businesses

	return response, nil
}
