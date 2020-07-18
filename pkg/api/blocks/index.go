package blocks

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services/database"
	"github.com/mewben/realty278/pkg/utils"
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
func Routes(g *fiber.Group, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Post("/blocks", func(c *fiber.Ctx) {
		log.Println("invoices.post")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		payload := fiber.Map{}
		if err := c.BodyParser(&payload); err != nil {
			log.Println("errrbodyparser", err)
			c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
			return
		}

		response, err := h.Create(payload)
		log.Println("aferecreate")
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(201).JSON(response)
	})
}
