package lots

import (
	"log"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Create lot
func (h *Handler) Create(data *Payload) (*models.LotModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		log.Println("error validate", err)
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	// get project if exists
	foundProject := h.DB.FindByID(h.Ctx, enums.CollProjects, data.ProjectID.Hex(), h.Business.ID)
	if foundProject == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFoundProject)
	}

	lot := models.NewLotModel(h.User.ID, h.Business.ID)
	lot.ProjectID = data.ProjectID
	lot.Name = data.Name
	lot.Area = data.Area
	lot.Price = data.Price
	lot.PriceAddon = data.PriceAddon
	lot.MetaModel = data.MetaModel

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollLots, lot)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	// TODO: create hooks
	lot = doc.(*models.LotModel)

	return lot, nil
}
