package lots

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Get lots
func (h *Handler) Get(projectID string) (*ResponseList, error) {
	log.Println("--get--projectID", projectID)
	filter := bson.D{
		{
			Key:   "businessID",
			Value: h.Business.ID,
		},
	}

	if projectID != "" {
		if projectID == "null" {
			log.Println("null projectid")
			filter = append(filter, bson.E{
				Key:   "projectID",
				Value: nil,
			})
		} else {
			log.Println("-get projectoid")
			projectOID, err := primitive.ObjectIDFromHex(projectID)
			if err != nil {
				return nil, errors.NewHTTPError(errors.ErrInputInvalid)
			}
			log.Println("---projectOID", projectOID)
			filter = append(filter, bson.E{
				Key:   "projectID",
				Value: projectOID,
			})
		}
	}

	opts := options.Find().SetSort(bson.D{
		{
			Key:   "createdAt",
			Value: -1,
		},
	})
	cursor, err := h.DB.Find(h.Ctx, enums.CollLots, filter, opts)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	lots := make([]*models.LotModel, 0)
	if err = cursor.All(h.Ctx, &lots); err != nil {
		log.Println("err cursor", err)
		return nil, err
	}

	response := &ResponseList{
		Total: len(lots), // do this for now
		Data:  lots,
	}

	return response, nil
}
