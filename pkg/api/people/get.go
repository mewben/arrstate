package people

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Get people
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
	cursor, err := h.DB.Find(h.Ctx, enums.CollPeople, filter, opts)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	people := make([]*models.PersonModel, 0)
	if err = cursor.All(h.Ctx, &people); err != nil {
		log.Println("err cursor", err)
		return nil, err
	}

	response := &ResponseList{
		Total: len(people), // do this for now
		Data:  people,
	}

	return response, nil
}
