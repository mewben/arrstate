package blocks

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateDefaultEntityBlocks -
func (h *Handler) CreateDefaultEntityBlocks(entityType string, entityID primitive.ObjectID, blocks []fiber.Map) error {
	for _, block := range blocks {
		block["entityType"] = entityType
		block["entityID"] = entityID
		if _, err := h.Create(block); err != nil {
			log.Println("err createblock", entityType, block["type"], block["title"])
			return err
		}
	}
	return nil
}
