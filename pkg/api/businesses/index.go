package businesses

import (
	"context"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"

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
	models.BusinessModel
}

// ResponseList -
type ResponseList struct {
	Total int                     `json:"total"`
	Data  []*models.BusinessModel `json:"list"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(g fiber.Router, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Get("/dashboard", func(c *fiber.Ctx) {
		log.Println("businesses.get.dashboard")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.GenerateDashboard()
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)
	})

	g.Get("/businesses", func(c *fiber.Ctx) {
		log.Println("businesses.get")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.Get()
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)
	})

	g.Get("/businesses/countries", func(c *fiber.Ctx) {
		log.Println("businesses.get.countries")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.GetCountries()
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})

	g.Get("/businesses/current", func(c *fiber.Ctx) {
		log.Println("businesses.get.countries")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		c.Status(200).JSON(h.Business)
	})

	g.Get("/businesses/currencies", func(c *fiber.Ctx) {
		log.Println("businesses.get.currencies")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.GetCurrencies()
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})

	g.Post("/businesses", func(c *fiber.Ctx) {
		log.Println("businesses.post")
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
}
