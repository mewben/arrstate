package properties

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// GetOne property
func (h *Handler) GetOne(id string) (*models.PropertyModel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	doc, err := h.DB.FindByID(h.Ctx, enums.CollProperties, oid, h.Business.ID)
	if err != nil {
		return nil, err
	}
	return doc.(*models.PropertyModel), nil

}
