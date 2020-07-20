package properties

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/api/invoices"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/utils"
)

// AcquireHook -
// generate invoices
func (h *Handler) AcquireHook(acquisition *AcquisitionPayload, property *models.PropertyModel) error {
	log.Println("acquire hoook")
	invoicesPayload := make([]*invoices.Payload, 0)
	issueDate := time.Now()

	if acquisition.PaymentScheme == enums.PaymentSchemeCash {
		dueDate := utils.EndOfMonth(issueDate)
		// create only one invoice
		invoicesPayload = append(invoicesPayload, &invoices.Payload{
			Name:    fmt.Sprintf("%s - 0", property.Name),
			DueDate: &dueDate,
			Blocks: []fiber.Map{
				{
					"type":   enums.InvoiceBlockItem,
					"title":  property.Name + " - Cash",
					"amount": property.Price,
				},
			},
		})
	} else if acquisition.PaymentScheme == enums.PaymentSchemeInstallment {
		// downpayment
		dueDates := make([]time.Time, acquisition.Terms+1)

		dueDates[0] = utils.EndOfMonth(issueDate)
		dueDate := dueDates[0]

		invoicesPayload = append(invoicesPayload, &invoices.Payload{
			Name:    fmt.Sprintf("%s - 0/%d", property.Name, acquisition.Terms),
			DueDate: &dueDates[0],
			Blocks: []fiber.Map{
				{
					"type":   enums.InvoiceBlockItem,
					"title":  property.Name + " - Down Payment",
					"amount": acquisition.DownPayment,
				},
			},
		})
		// (totalPrice - downPayment) / terms
		var recurringAmount int64
		// recurringAmount = int64(math.Round(float64(property.Price-acquisition.DownPayment)) / float64(acquisition.Terms))
		recurringAmount = int64(math.Round(float64(property.Price-acquisition.DownPayment) / float64(acquisition.Terms)))
		// 12 invoices if terms = 12
		for i := 1; i <= acquisition.Terms; i++ {
			dueDates[i] = dueDate.AddDate(0, 0, 1)
			dueDates[i] = utils.EndOfMonth(dueDates[i])
			dueDate = dueDates[i]
			log.Println("dueDate:", dueDate)
			invoicesPayload = append(invoicesPayload, &invoices.Payload{
				Name:    fmt.Sprintf("%s - %d/%d", property.Name, i, acquisition.Terms),
				DueDate: &dueDates[i],
				Blocks: []fiber.Map{
					{
						"type":   enums.InvoiceBlockItem,
						"title":  fmt.Sprintf("%s %d/%d", property.Name, i, acquisition.Terms),
						"amount": recurringAmount,
					},
				},
			})
		}
		log.Println("dueDates:", dueDates)
	}

	if len(invoicesPayload) > 0 {
		log.Println("before create invoices")
		for _, payload := range invoicesPayload {
			// get business every loop to get the updated next invoice Seq
			// get business
			business := h.DB.FindBusiness(h.Ctx, h.Business.ID)
			if business == nil {
				return errors.NewHTTPError(errors.ErrNotFoundBusiness)
			}

			h := invoices.Handler{
				Ctx:      h.Ctx,
				DB:       h.DB,
				User:     h.User,
				Business: business,
			}
			payload.To = &models.FromToModel{
				ID:         acquisition.ClientID,
				EntityType: enums.EntityPerson,
			}
			payload.Status = enums.StatusPending
			payload.ProjectID = property.ProjectID
			payload.PropertyID = &property.ID
			payload.IssueDate = &issueDate
			if _, err := h.Create(payload); err != nil {
				// TODO: handle rollback
				return err
			}
		}

	}

	return nil

}
