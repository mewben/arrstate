package blocks

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
)

// Get blocks
func (h *Handler) Get(data *GetPayload) (*ResponseList, error) {
	response := &ResponseList{}
	if len(data.IDs) == 0 {
		return response, nil
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: bson.M{"$in": data.IDs},
		},
		{
			Key:   "entityType",
			Value: data.EntityType,
		},
		{
			Key:   "entityID",
			Value: data.EntityID,
		},
		{
			Key:   "businessID",
			Value: h.Business.ID,
		},
	}
	cursor, err := h.DB.Find(h.Ctx, enums.CollBlocks, filter)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}
	blocks := make([]fiber.Map, 0)
	if err = cursor.All(h.Ctx, &blocks); err != nil {
		log.Println("err cursor", err)
		return nil, err
	}

	response.Data = blocks
	response.Total = len(blocks)

	return response, nil

}
