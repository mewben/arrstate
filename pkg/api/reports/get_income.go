package reports

import (
	"log"
	"time"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetIncome -
// date format = MM-DD-YYYY
func (h *Handler) GetIncome(from, to string) (*ResponseList, error) {
	// parse from and to to dates
	layout := "01-02-2006"
	fromDate, err := time.Parse(layout, from)
	if err != nil {
		return nil, err
	}
	toDate, err := time.Parse(layout, to)
	if err != nil {
		return nil, err
	}
	// add 1 day to include
	toDate = toDate.Add(24 * 60 * 60 * time.Second)

	log.Println("from", fromDate)
	log.Println("to", toDate)

	filter := bson.D{
		{
			Key:   "businessID",
			Value: h.Business.ID,
		},
		{
			Key: "paidAt",
			Value: bson.M{
				"$gte": fromDate,
				"$lt":  toDate,
			},
		},
	}
	opts := options.Find().SetLimit(20).SetSort(bson.M{"paidAt": -1})

	total, err := h.DB.Count(h.Ctx, enums.CollInvoices, filter)
	if err != nil {
		return nil, err
	}

	cursor, err := h.DB.Find(h.Ctx, enums.CollInvoices, filter, opts)
	if err != nil {
		return nil, err
	}

	invoices := make([]*models.InvoiceModel, 0)
	if err = cursor.All(h.Ctx, &invoices); err != nil {
		log.Println("err cursor", err)
		return nil, err
	}

	// data results is limited to max 20 to minimize resources.
	response := &ResponseList{
		Total: int(total),
		Data:  invoices,
	}

	return response, nil
}
