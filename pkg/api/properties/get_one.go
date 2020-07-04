package properties

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// GetOne property
func (h *Handler) GetOne(id string) (*models.PropertyModel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	doc := h.DB.FindByID(h.Ctx, enums.CollProperties, oid, h.Business.ID)
	if doc == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	return doc.(*models.PropertyModel), nil

}
