package invoices

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// GetOne invoice
func (h *Handler) GetOne(id string) (*models.InvoiceModel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	doc, err := h.DB.FindByID(h.Ctx, enums.CollInvoices, oid, h.Business.ID)
	if err != nil {
		return nil, err
	}

	return doc.(*models.InvoiceModel), nil

}
