package projects

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// Get projects
func (h *Handler) Get() (*ResponseList, error) {
	filter := bson.D{
		{
			Key:   "businessID",
			Value: h.Business.ID,
		},
	}
	opts := options.Find().SetSort(bson.D{
		{
			Key:   "createdAt",
			Value: -1,
		},
	})
	cursor, err := h.DB.Find(h.Ctx, enums.CollProjects, filter, opts)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	projects := make([]*models.ProjectModel, 0)
	if err = cursor.All(h.Ctx, &projects); err != nil {
		log.Println("err cursor", err)
		return nil, err
	}

	response := &ResponseList{
		Total: len(projects), // do this for now
		Data:  projects,
	}

	return response, nil
}
