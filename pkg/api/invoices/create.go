package invoices

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/api/blocks"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/utils"
)

// Create invoice
func (h *Handler) Create(data *Payload) (*models.InvoiceModel, error) {
	invoice := models.NewInvoiceModel(h.User.ID, h.Business.ID)
	// Validate payload
	// 1. check from
	if data.From != nil {
		if data.From.EntityType == enums.EntityPerson {
			_, err := h.DB.FindByID(h.Ctx, enums.CollPeople, *data.From.ID, h.Business.ID)
			if err != nil {
				return nil, err
			}
			invoice.From = data.From
		} else {
			return nil, errors.NewHTTPError(errors.ErrInputInvalid)
		}
	} else {
		invoice.From = &models.FromToModel{
			ID:         &h.Business.ID,
			EntityType: enums.EntityBusiness,
		}
	}
	// 2. check toID
	if data.To != nil {
		if data.To.EntityType == enums.EntityPerson {
			_, err := h.DB.FindByID(h.Ctx, enums.CollPeople, *data.To.ID, h.Business.ID)
			if err != nil {
				return nil, err
			}
			invoice.To = data.To
		} else {
			return nil, errors.NewHTTPError(errors.ErrInputInvalid)
		}
	}
	// 3. check projectID
	if data.ProjectID != nil {
		_, err := h.DB.FindByID(h.Ctx, enums.CollProjects, *data.ProjectID, h.Business.ID)
		if err != nil {
			return nil, err
		}
		invoice.ProjectID = data.ProjectID
	}
	// 4. check propertyID
	if data.PropertyID != nil {
		_, err := h.DB.FindByID(h.Ctx, enums.CollProperties, *data.PropertyID, h.Business.ID)
		if err != nil {
			return nil, err
		}
		invoice.PropertyID = data.PropertyID
	}
	// 5. TODO: check discount
	// 6. check tax
	if data.Tax < 0 {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// Validate status
	invoice.Status = data.Status
	if data.Status == "" {
		invoice.Status = enums.StatusDraft
	}
	if !utils.Contains(allowedStatuses, invoice.Status) {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// assign to invoice model
	invoice.No = h.Business.Invoices.NextSeq
	invoice.Name = data.Name
	if invoice.Name == "" {
		invoice.Name = strconv.Itoa(h.Business.Invoices.NextSeq)
	}
	invoice.IssueDate = data.IssueDate
	invoice.DueDate = data.DueDate
	invoice.Tax = data.Tax
	invoice.Discount = data.Discount
	invoice.Blocks = make([]primitive.ObjectID, 0)
	// TODO: amounts will be calculated on the blocks hooks
	invoice.Total = 0
	invoice.SubTotal = 0

	// prepare blocks
	if len(data.Blocks) == 0 {
		data.Blocks = []fiber.Map{
			{
				"type": enums.InvoiceBlockIntro,
			},
			{
				"type": enums.InvoiceBlockItem,
			},
			{
				"type": enums.InvoiceBlockSummary,
			},
		}
	} else {
		// append intro and summary
		data.Blocks = append([]fiber.Map{{"type": enums.InvoiceBlockIntro}}, data.Blocks...)
		data.Blocks = append(data.Blocks, fiber.Map{"type": enums.InvoiceBlockSummary})
	}

	// insert invoice
	doc, err := h.DB.InsertOne(h.Ctx, enums.CollInvoices, invoice)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}
	invoice = doc.(*models.InvoiceModel)

	blocksHandler := blocks.Handler{
		DB:       h.DB,
		Ctx:      h.Ctx,
		User:     h.User,
		Business: h.Business,
	}
	if err := blocksHandler.CreateDefaultEntityBlocks(enums.EntityInvoice, invoice.ID, data.Blocks); err != nil {
		log.Println("err createdefaultentityblocks", err)
		return nil, err
	}

	// AfterCreate hook
	if err := h.CreateHook(invoice); err != nil {
		return nil, err
	}

	// get fresh invoice in case it has been updated by creating the blocks
	return h.GetOne(invoice.ID.Hex())

}
