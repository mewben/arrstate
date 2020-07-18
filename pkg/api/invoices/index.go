package invoices

import (
	"context"
	"log"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services/database"
	"github.com/mewben/realty278/pkg/utils"
)

// Handler -
type Handler struct {
	DB       *database.Service
	Ctx      context.Context
	User     *models.UserModel
	Business *models.BusinessModel
}

// Payload -
type Payload struct {
	From       *models.FromToModel `json:"from"`
	To         *models.FromToModel `json:"to"`
	ProjectID  *primitive.ObjectID `json:"projectID"`
	PropertyID *primitive.ObjectID `json:"propertyID"`
	IssueDate  *time.Time          `json:"issueDate"`
	DueDate    *time.Time          `json:"dueDate"`
	Blocks     []fiber.Map         `json:"blocks"`
	Discount   string              `json:"discount"`
	Tax        float64             `json:"tax"`
}

// ResponseList -
type ResponseList struct {
	Total int                    `json:"total"`
	Data  []*models.InvoiceModel `json:"list"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(g *fiber.Group, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Get("/invoices", func(c *fiber.Ctx) {
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.Get(c.Query("propertyID"))
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)
	})

	g.Get("/invoices/:invoiceID", func(c *fiber.Ctx) {
		log.Println("invoices.getOne")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.GetOne(c.Params("invoiceID"))
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)
	})

	g.Post("/invoices", func(c *fiber.Ctx) {
		log.Println("invoices.post")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		payload := &Payload{}
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
