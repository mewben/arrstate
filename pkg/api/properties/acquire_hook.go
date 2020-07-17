package properties

import (
	"log"

	"github.com/mewben/realty278/pkg/models"
)

// AcquireHook -
// generate invoices
func (h *Handler) AcquireHook(property *models.PropertyModel, downPayment float64) error {
	log.Println("acquire hoook")
	// acquisition := property.Acquisition
	// invoices := make([]interface{}, 0)

	// if data.PaymentScheme == enums.PaymentSchemeCash {
	// 	invoice := models.NewInvoiceModel(h.User.ID, h.Business.ID)
	// 	invoice.ID = primitive.NewObjectID()
	// 	invoice.Total = downPayment

	// 	// create invoice items block

	// 	invoices = append(invoices, &models.InvoiceModel{

	// 	})
	// }

	// if (len(invoices) > 0) {
	// 	_, err := h.DB.InsertMany(h.Ctx, enums.CollInvoices, invoices)
	// 	return err
	// }

	return nil

}
