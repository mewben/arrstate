package auth

import (
	"context"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/services/database"
	"github.com/mewben/realty278/pkg/utils"
)

// Handler -
type Handler struct {
	DB  *database.Service
	Ctx context.Context
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(app *fiber.App, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	app.Post("/auth/signup", func(c *fiber.Ctx) {
		h.Ctx = c.Fasthttp
		payload := &SignupPayload{}

		if err := c.BodyParser(&payload); err != nil {
			c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
			return
		}

		response, err := h.Signup(payload)
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(201).JSON(response)
	})

	app.Post("/auth/signin", func(c *fiber.Ctx) {
		h.Ctx = c.Fasthttp
		payload := &SigninPayload{}

		if err := c.BodyParser(&payload); err != nil {
			c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
			return
		}

		// get business domain
		domain := utils.GetSubdomain(string(c.Fasthttp.Request.Header.Peek("origin")))

		response, err := h.Signin(payload, domain)
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})
}
