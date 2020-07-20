package invoices

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/models"
)

// CreateHook -
func (h *Handler) CreateHook(invoice *models.InvoiceModel) error {
	// 1. increase business.invoices.nextSeq
	upd := bson.D{
		{
			Key:   "$inc",
			Value: bson.M{"invoices.nextSeq": 1},
		},
	}
	_, err := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollBusinesses, h.Business.ID, upd)
	if err != nil {
		return err
	}

	return nil
}
