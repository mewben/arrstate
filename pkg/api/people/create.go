package people

import (
	"encoding/json"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/models"
)

// Payload -
type Payload struct {
	UserID     string `json:"userID"`
	BusinessID string `json:"businessID"`
	Role       string `json:"role"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

// Create User
func (h *Handler) Create(data *Payload) (*models.PersonModel, error) {
	person := models.NewPersonModel()
	person.UserID = data.UserID
	person.BusinessID = data.BusinessID
	person.Role = data.Role
	person.GivenName = data.GivenName
	person.FamilyName = data.FamilyName
	person.Country = "PH"

	result, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	log.Println(string(result))

	insertResult, err := h.DB.InsertOne(h.Ctx, enums.CollPeople, person)
	if err != nil {
		return nil, err
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: insertResult.InsertedID,
		},
	}
	insertedModel := h.DB.FindOne(h.Ctx, enums.CollPeople, filter)
	if insertedModel == nil {
		return nil, errors.New("Not found")
	}

	return insertedModel.(*models.PersonModel), nil
}
