package invoices

import (
	"context"
	"log"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services/database"
	"github.com/mewben/arrstate/pkg/utils"
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
	Name       string                  `json:"name"`
	Status     string                  `json:"status"`
	From       *models.FromToModel     `json:"from"`
	To         *models.FromToModel     `json:"to"`
	ProjectID  *primitive.ObjectID     `json:"projectID"`
	PropertyID *primitive.ObjectID     `json:"propertyID"`
	IssueDate  *time.Time              `json:"issueDate"`
	DueDate    *time.Time              `json:"dueDate"`
	Blocks     []fiber.Map             `json:"blocks"`
	AddOrLess  []models.AddOrLessModel `json:"addOrLess"`
	// Discount   string              `json:"discount"`
	// Tax        int64               `json:"tax"`
}

// PaymentPayload -
type PaymentPayload struct {
	ReceiptNo string             `json:"receiptNo" validate:"required"`
	InvoiceID primitive.ObjectID `json:"invoiceID" validate:"required"`
}

// ResponseList -
type ResponseList struct {
	Total int                    `json:"total"`
	Data  []*models.InvoiceModel `json:"list"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()
var allowedStatuses []string

func init() {
	allowedStatuses = []string{enums.StatusDraft, enums.StatusPending, enums.StatusOverdue, enums.StatusPaid}
}

// Routes -
func Routes(g fiber.Router, db *mongo.Database) {
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

		propertyID := c.Query("propertyID")
		status := c.Query("status")

		response, err := h.Get(propertyID, status)
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
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(201).JSON(response)

	})

	g.Post("/invoices/pay", func(c *fiber.Ctx) {
		log.Println("invoices.pay")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		payload := &PaymentPayload{}
		if err := c.BodyParser(&payload); err != nil {
			log.Println("errrbodyparser", err)
			c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
			return
		}
		utils.PrettyJSON(payload)

		response, err := h.Pay(payload)
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})

}
