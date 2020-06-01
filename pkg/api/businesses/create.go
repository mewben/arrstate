package businesses

import (
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Payload -
type Payload struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

// Validate Payload
func (v Payload) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(
			&v.Name,
			validation.Required,
		),
		validation.Field(
			&v.Domain,
			validation.Required,
			validation.Length(3, 254),
		),
	)
}

// Create business
func (h *Handler) Create(data *Payload) (*models.BusinessModel, error) {
	// validate payload
	if err := data.Validate(); err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	cleanedDomain := slug.Make(data.Domain)

	// check duplicate domain
	filter := bson.D{
		{
			Key:   "domain",
			Value: cleanedDomain,
		},
	}
	businessFound := h.DB.FindOne(h.Ctx, enums.CollBusinesses, filter)
	if businessFound != nil {
		return nil, errors.NewHTTPError(errors.ErrDomainDuplicate)
	}

	// check pass
	business := models.NewBusinessModel()
	business.Name = data.Name
	business.Domain = cleanedDomain
	insertResult, err := h.DB.InsertOne(h.Ctx, enums.CollBusinesses, business)
	if err != nil {
		log.Println("insertonerr", err)
		return nil, err
	}

	filter = bson.D{
		{
			Key:   "_id",
			Value: insertResult.InsertedID,
		},
	}
	insertedModel := h.DB.FindOne(h.Ctx, enums.CollBusinesses, filter)
	if insertedModel == nil {
		log.Println("findoneerr")
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}

	return insertedModel.(*models.BusinessModel), nil
}
