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
func (h *Handler) Get(propertyID, status string) (*ResponseList, error) {

	filter := bson.D{
		{
			Key:   "businessID",
			Value: h.Business.ID,
		},
	}

	if propertyID != "" {
		propertyOID, err := primitive.ObjectIDFromHex(propertyID)
		if err != nil {
			log.Println("err: propertyID")
			return nil, errors.NewHTTPError(errors.ErrInputInvalid)
		}

		filter = append(filter, bson.E{
			Key:   "propertyID",
			Value: propertyOID,
		})
	}
	if status != "" {
		filter = append(filter, bson.E{
			Key:   "status",
			Value: status,
		})
	}

	log.Println("--filter:", filter)

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
