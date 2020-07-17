package invoices

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Get invoices
func (h *Handler) Get(propertyID string) (*ResponseList, error) {
	propertyOID, err := primitive.ObjectIDFromHex(propertyID)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	filter := bson.D{
		{
			Key:   "propertyID",
			Value: propertyOID,
		},
	}
	cursor, err := h.DB.Find(h.Ctx, enums.CollInvoices, filter)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	invoices := make([]*models.InvoiceModel, 0)
	if err = cursor.All(h.Ctx, &invoices); err != nil {
		log.Println("err cursor", err)
		return nil, err
	}

	response := &ResponseList{
		Total: len(invoices), // do this for now
		Data:  invoices,
	}

	return response, nil
}
