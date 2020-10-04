package files

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/utils"
)

// Get files
func (h *Handler) Get(params *Params) (*ResponseList, error) {

	log.Println("files.get", params)
	filter := bson.D{
		{
			Key:   "businessID",
			Value: h.Business.ID,
		},
	}

	if params.EntityType != nil {
		filter = append(filter, bson.E{
			Key:   "entityType",
			Value: params.EntityType,
		})
	}

	if params.EntityID != nil {
		filter = append(filter, bson.E{
			Key:   "entityID",
			Value: params.EntityID.ObjectID,
		})
	}

	opts := options.Find().SetSort(bson.D{
		{
			Key:   "createdAt",
			Value: -1,
		},
	})

	utils.PrettyJSON(filter)

	cursor, err := h.DB.Find(h.Ctx, enums.CollFiles, filter, opts)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	files := make([]*models.FileModel, 0)
	if err = cursor.All(h.Ctx, &files); err != nil {
		log.Println("err cursor", err)
		return nil, err
	}

	response := &ResponseList{
		Total: len(files), // do this for now
		Data:  files,
	}

	return response, nil
}
