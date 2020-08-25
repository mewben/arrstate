package people

import (
	"context"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
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
	Person   *models.PersonModel
}

// Payload -
type Payload struct {
	models.PersonModel
}

// ResponseList -
type ResponseList struct {
	Total int                   `json:"total"`
	Data  []*models.PersonModel `json:"list"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(g fiber.Router, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Get("/people", func(c *fiber.Ctx) {
		log.Println("people.get")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		var role []string
		r := c.Fasthttp.QueryArgs().PeekMulti("role")
		for _, rr := range r {
			role = append(role, utils.GetString(rr))
		}

		log.Println("--role", role)

		response, err := h.Get(role)
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)
	})

	g.Get("/people/:personID", func(c *fiber.Ctx) {
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.GetOne(c.Params("personID"))
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})

	g.Post("/people", func(c *fiber.Ctx) {
		log.Println("people.post")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
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

		if utils.Contains(payload.Role, enums.RoleOwner) {
			c.Status(400).JSON(errors.NewHTTPError(errors.ErrOwner))
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

	g.Put("/people", func(c *fiber.Ctx) {
		log.Println("people.put")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
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

		response, err := h.Edit(payload)
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})

	g.Put("/people/:personID/locale", func(c *fiber.Ctx) {
		log.Println("people.edit.locale")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, h.Person, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		payload := &models.Locale{}
		if err := c.BodyParser(&payload); err != nil {
			log.Println("errrbodyparser", err)
			c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
			return
		}

		response, err := h.EditLocale(c.Params("personID"), payload)
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})

	g.Delete("/people/:personID", func(c *fiber.Ctx) {
		log.Println("people.delete")
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.Remove(c.Params("personID"))
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})
}
