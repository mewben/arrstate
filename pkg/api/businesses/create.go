package businesses

import (
	"log"

	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Create business
func (h *Handler) Create(data *Payload) (*models.BusinessModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	cleanedDomain := slug.Make(data.Domain)

	// check duplicate domain
	filter := bson.D{
		{
			Key:   "domain",
			Value: cleanedDomain,
		},
	}
	businessFound, _ := h.DB.FindOne(h.Ctx, enums.CollBusinesses, filter)
	log.Println("---businessFound:", businessFound)
	if businessFound != nil {
		return nil, errors.NewHTTPError(errors.ErrDomainDuplicate)
	}

	// check pass
	business := models.NewBusinessModel()
	business.Name = data.Name
	business.Domain = cleanedDomain
	doc, err := h.DB.InsertOne(h.Ctx, enums.CollBusinesses, business)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	business = doc.(*models.BusinessModel)

	return business, nil
}
