package businesses

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/api/users"
	"github.com/mewben/arrstate/pkg/models"
)

// GenerateDashboard -
func (h *Handler) GenerateDashboard() (*models.MeModel, error) {
	var err error
	// TODO: optimize use waitGroup
	dashData := h.Business.Dashboard
	if len(dashData) == 0 {
		dashData = make(map[string]models.DashboardModel, 0)
	}
	// generate dashboard data
	dashData["projects"], err = h.query(enums.CollProjects, dashData)
	if err != nil {
		return nil, err
	}

	dashData["properties"], err = h.query(enums.CollProperties, dashData)
	if err != nil {
		return nil, err
	}

	dashData["people"], err = h.query(enums.CollPeople, dashData)
	if err != nil {
		return nil, err
	}

	// update business
	upd := bson.D{
		{
			Key: "$set",
			Value: fiber.Map{
				"dashboard": dashData,
			},
		},
	}
	_, err = h.DB.FindByIDAndUpdate(h.Ctx, enums.CollBusinesses, h.Business.ID, upd)
	if err != nil {
		return nil, err
	}

	userHandler := users.Handler{
		DB:       h.DB,
		Ctx:      h.Ctx,
		User:     h.User,
		Business: h.Business,
	}
	return userHandler.GetMe()
}

func (h *Handler) query(collectionName string, dashData map[string]models.DashboardModel) (models.DashboardModel, error) {
	dashModel := models.DashboardModel{}
	filter := bson.D{{
		Key:   "businessID",
		Value: h.Business.ID,
	}}
	count, err := h.DB.Count(h.Ctx, collectionName, filter)
	if err != nil {
		return dashModel, err
	}

	var m string
	var label string
	switch collectionName {
	case enums.CollProjects:
		m = "projects"
		label = "Projects"
		break
	case enums.CollProperties:
		m = "properties"
		label = "Properties"
		break
	case enums.CollPeople:
		m = "people"
		label = "People"
		break
	}
	if dashData[m].Label == "" {
		dashModel.Label = label
	} else {
		dashModel.Label = dashData[m].Label
	}
	dashModel.Total = float64(count)

	return dashModel, nil
}
