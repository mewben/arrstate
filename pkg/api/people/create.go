package people

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

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
	personFound, _ := h.DB.FindOne(h.Ctx, enums.CollPeople, filter)
	if personFound != nil {
		return nil, errors.NewHTTPError(errors.ErrDuplicateDomain)
	}

	person := models.NewPersonModel(h.User.ID, h.Business.ID)
	person.UserID = data.UserID
	person.Email = data.Email
	person.Role = data.Role
	person.Name = data.Name
	person.Address = data.Address
	person.MetaModel = data.MetaModel
	person.CommissionPerc = data.CommissionPerc
	person.CustomFields = data.CustomFields
	if person.Locale.Language == "" {
		person.Locale.Language = enums.DefaultLanguage
	}
	if person.Locale.DateFormat == "" {
		person.Locale.DateFormat = enums.DefaultDateFormat
	}
	if person.Locale.TimeFormat == "" {
		person.Locale.TimeFormat = enums.DefaultTimeFormat
	}
	if person.Locale.TimestampFormat == "" {
		person.Locale.TimestampFormat = enums.DefaultTimestampFormat
	}
	if person.Locale.WeekStartDay == "" {
		person.Locale.WeekStartDay = "sunday"
	}

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollPeople, person)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	person = doc.(*models.PersonModel)

	return person, nil
}
