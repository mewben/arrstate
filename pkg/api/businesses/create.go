package businesses

import (
	"errors"
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
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
		),
	)
}

// Create business
func (h *Handler) Create(data *Payload) (*models.BusinessModel, error) {

	business := models.NewBusinessModel()
	business.Name = data.Name
	business.Domain = data.Domain

	insertResult, err := h.DB.InsertOne(h.Ctx, enums.CollBusinesses, business)
	if err != nil {
		log.Println("insertonerr", err)
		return nil, err
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: insertResult.InsertedID,
		},
	}
	insertedModel := h.DB.FindOne(h.Ctx, enums.CollBusinesses, filter)
	if insertedModel == nil {
		log.Println("findoneerr")
		return nil, errors.New("Not found")
	}

	return insertedModel.(*models.BusinessModel), nil
}
