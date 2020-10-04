package invoices

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// Pay invoice
func (h *Handler) Pay(data *PaymentPayload) (*models.InvoiceModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	// get current invoice
	foundInvoice, err := h.DB.FindByID(h.Ctx, enums.CollInvoices, data.InvoiceID, h.Business.ID)
	if err != nil {
		return nil, err
	}
	invoice := foundInvoice.(*models.InvoiceModel)

	if invoice.Status == enums.StatusPaid {
		return nil, errors.NewHTTPError(errors.ErrInvoiceAlreadyPaid)
	}

	// check receiptNo for duplicate
	filter := bson.D{
		{
			Key:   "receiptNo",
			Value: data.ReceiptNo,
		},
		{
			Key:   "businessID",
			Value: h.Business.ID,
		},
	}
	if _, err = h.DB.FindOne(h.Ctx, enums.CollInvoices, filter); err == nil {
		return nil, errors.NewHTTPError(errors.ErrDuplicateReceipt)
	}

	// pass
	upd := fiber.Map{
		"status":    enums.StatusPaid,
		"receiptNo": data.ReceiptNo,
	}
	op := bson.D{
		{
			Key:   "$set",
			Value: upd,
		},
		{
			Key: "$currentDate",
			Value: fiber.Map{
				"paidAt": true,
			},
		},
	}
	doc, err := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollInvoices, invoice.ID, op)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate, err)
	}

	invoice = doc.(*models.InvoiceModel)
	return invoice, nil

}
