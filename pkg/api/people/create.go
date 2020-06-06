package people

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Payload2 -
type Payload2 struct {
	UserID     primitive.ObjectID `json:"userID"`
	BusinessID primitive.ObjectID `json:"businessID" validate:"required"`
	Role       string             `json:"role" validate:"required"`
	GivenName  string             `json:"givenName" validate:"required"`
	FamilyName string             `json:"familyName"`
}

// Create Person
func (h *Handler) Create(data *Payload) (*models.PersonModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	// check duplicate person in this business
	filter := bson.D{
		{
			Key:   "email",
			Value: data.Email,
		},
		{
			Key:   "businessID",
			Value: h.Business.ID,
		},
	}
	personFound := h.DB.FindOne(h.Ctx, enums.CollPeople, filter)
	if personFound != nil {
		return nil, errors.NewHTTPError(errors.ErrDomainDuplicate)
	}

	person := models.NewPersonModel(h.User.ID, h.Business.ID)
	person.UserID = data.UserID
	person.Email = data.Email
	person.Role = data.Role
	person.GivenName = data.GivenName
	person.FamilyName = data.FamilyName
	person.Address = data.Address
	person.MetaModel = data.MetaModel
	person.CommissionPerc = data.CommissionPerc
	person.CustomFields = data.CustomFields

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollPeople, person)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	person = doc.(*models.PersonModel)

	return person, nil
}
