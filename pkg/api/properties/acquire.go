package properties

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/utils"
)

// AcquisitionPayload -
type AcquisitionPayload struct {
	PropertyID *primitive.ObjectID `json:"propertyID"`
	models.AcquisitionModel
	DownPayment int64 `json:"downPayment"`
}

// Acquire property
func (h *Handler) Acquire(data *AcquisitionPayload) (*models.PropertyModel, error) {
	// 1. Validate Property
	if data.PropertyID == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFoundProperty)
	}
	foundProperty, err := h.DB.FindByID(h.Ctx, enums.CollProperties, *data.PropertyID, h.Business.ID)
	if err != nil {
		return nil, err
	}
	property := foundProperty.(*models.PropertyModel)
	if property.Status != enums.StatusAvailable {
		return nil, errors.NewHTTPError(errors.ErrPropertyAlreadyTaken)
	}

	// 2. validate Client
	if data.ClientID == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFoundPerson)
	}
	_, err = h.DB.FindByID(h.Ctx, enums.CollPeople, *data.ClientID, h.Business.ID)
	if err != nil {
		return nil, err
	}

	// 3. validate Agent
	if data.AgentID != nil {
		_, err = h.DB.FindByID(h.Ctx, enums.CollPeople, *data.AgentID, h.Business.ID)
		if err != nil {
			return nil, err
		}
	}

	// 4. validate paymentScheme
	if !utils.Contains(allowedPaymentSchemes, data.PaymentScheme) {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// prepare acquisition model
	status := enums.StatusOnGoing
	updatedAts := fiber.Map{
		"updatedAt":              "$$NOW",
		"acquisition.acquiredAt": "$$NOW",
	}
	if data.PaymentScheme == enums.PaymentSchemeInstallment {
		// validate further
		// 5. validate paymentPeriod
		if !utils.Contains(allowedPaymentPeriods, data.PaymentPeriod) {
			return nil, errors.NewHTTPError(errors.ErrInputInvalid)
		}
		// 6. validate terms
		if data.Terms <= 0 {
			return nil, errors.NewHTTPError(errors.ErrInputInvalid)
		}
		// 7. validate downpayment
		if data.DownPayment <= 0 {
			return nil, errors.NewHTTPError(errors.ErrInputInvalid)
		}

	} else if data.PaymentScheme == enums.PaymentSchemeCash {
		updatedAts["acquisition.completedAt"] = "$$NOW"
		status = enums.StatusAcquired
	}

	upd := fiber.Map{
		"status":      status,
		"acquisition": data.AcquisitionModel,
	}

	// use update pipeline to utilize update $$NOW
	op := bson.A{
		bson.D{
			{
				Key:   "$set",
				Value: upd,
			},
		},
		bson.D{
			{
				Key:   "$set",
				Value: updatedAts,
			},
		},
	}
	doc, err := h.DB.FindByIDAndUpdate(h.Ctx, enums.CollProperties, property.ID, op)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrUpdate)
	}

	// hooks,
	// create invoices

	property = doc.(*models.PropertyModel)

	return property, nil
}
