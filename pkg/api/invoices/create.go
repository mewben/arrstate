package invoices

import (
	"log"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Create invoice
func (h *Handler) Create(data *Payload) (*models.InvoiceModel, error) {
	invoice := models.NewInvoiceModel(h.User.ID, h.Business.ID)
	// Validate payload
	// 1. check from
	if data.From.ID != nil {
		if data.From.EntityType == enums.EntityPerson {
			_, err := h.DB.FindByID(h.Ctx, enums.CollPeople, *data.From.ID, h.Business.ID)
			if err != nil {
				return nil, err
			}
			invoice.From = data.From
		} else {
			return nil, errors.NewHTTPError(errors.ErrInputInvalid)
		}
	}
	// 2. check toID
	if data.To.ID != nil {
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

	// assign to invoice model
	invoice.IssueDate = data.IssueDate
	invoice.DueDate = data.DueDate

	// insert invoice
	doc, err := h.DB.InsertOne(h.Ctx, enums.CollInvoices, invoice)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}
	invoice = doc.(*models.InvoiceModel)

	// TODO: create blocks

	// get fresh invoice in case it has been updated by creating the blocks
	return h.GetOne(invoice.ID.Hex())

}
