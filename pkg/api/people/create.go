package people

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Payload -
type Payload struct {
	UserID     primitive.ObjectID `json:"userID"`
	BusinessID primitive.ObjectID `json:"businessID" validate:"required"`
	Role       string             `json:"role" validate:"required"`
	GivenName  string             `json:"givenName" validate:"required"`
	FamilyName string             `json:"familyName"`
}

// Create User
func (h *Handler) Create(data *Payload) (*models.PersonModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	person := models.NewPersonModel(data.UserID, data.BusinessID)
	person.UserID = data.UserID
	person.Role = data.Role
	person.GivenName = data.GivenName
	person.FamilyName = data.FamilyName
	person.Country = enums.DefaultCountry

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollPeople, person)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	person = doc.(*models.PersonModel)

	return person, nil
}
