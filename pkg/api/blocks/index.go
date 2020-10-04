package blocks

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services/database"
	"github.com/mewben/arrstate/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler -
type Handler struct {
	DB       *database.Service
	Ctx      context.Context
	User     *models.UserModel
	Business *models.BusinessModel
}

// ResponseList -
type ResponseList struct {
	Total int         `json:"total"`
	Data  []fiber.Map `json:"list"`
}

// GetPayload -
type GetPayload struct {
	IDs        []primitive.ObjectID `json:"ids"`
	EntityType string               `json:"entityType"`
	EntityID   primitive.ObjectID   `json:"entityID"`
}

var allowedInvoiceBlocks []string

func init() {
	allowedInvoiceBlocks = []string{
		enums.InvoiceBlockIntro,
		enums.InvoiceBlockItem,
		enums.InvoiceBlockSummary,
	}
}

// NewBaseBlock -
func NewBaseBlock(arg ...primitive.ObjectID) *BaseBlock {
	return &BaseBlock{
		BlockModel: *models.NewBlockModel(arg...),
	}
}

// NewInvoiceItemBlock -
func NewInvoiceItemBlock(arg ...primitive.ObjectID) *InvoiceItemBlock {
	return &InvoiceItemBlock{
		InvoiceItemBlockModel: *models.NewInvoiceItemBlockModel(arg...),
	}
}

// Routes -
func Routes(g fiber.Router, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	// We use Post here because we supply a pretty large payload of blockIDs
	g.Post("/blocks/get", func(c *fiber.Ctx) error {
		log.Println("blocks.get")
		var err error
		h.Ctx = c.Context()
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			return c.Status(400).JSON(err)
		}

		payload := &GetPayload{}
		if err := c.BodyParser(&payload); err != nil {
			log.Println("errrbodyparser", err)
			return c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
		}

		response, err := h.Get(payload)
		if err != nil {
			log.Println("errrrrr", err)
			return c.Status(400).JSON(err)
		}
		return c.Status(200).JSON(response)

	})

	g.Post("/blocks", func(c *fiber.Ctx) error {
		log.Println("blocks.post")
		var err error
		h.Ctx = c.Context()
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			return c.Status(400).JSON(err)
		}

		payload := fiber.Map{}
		if err := c.BodyParser(&payload); err != nil {
			log.Println("errrbodyparser", err)
			return c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
		}

		response, err := h.Create(payload)
		if err != nil {
			log.Println("errrrrr", err)
			return c.Status(400).JSON(err)
		}
		return c.Status(201).JSON(response)
	})
}
