package auth

import (
	"context"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services/database"
	"github.com/mewben/arrstate/pkg/utils"
)

// Handler -
type Handler struct {
	DB  *database.Service
	Ctx context.Context
}

// SigninResponse -
type SigninResponse struct {
	Token string            `json:"token"`
	User  *models.UserModel `json:"user"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(app *fiber.App, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	app.Post("/auth/signup", func(c *fiber.Ctx) error {
		h.Ctx = c.Context()
		payload := &SignupPayload{}

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
		}

		response, err := h.Signup(payload)
		if err != nil {
			log.Println("errrrrr", err)
			return c.Status(400).JSON(err)
		}
		return c.Status(201).JSON(response)
	})

	app.Post("/auth/signin", func(c *fiber.Ctx) error {
		h.Ctx = c.Context()
		payload := &SigninPayload{}

		log.Println("signinb4parse")
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
		}
		log.Println("signinafterparse")

		// get business domain
		domain := utils.GetSubdomain(string(c.Context().Request.Header.Peek("origin")))

		response, err := h.Signin(payload, domain)
		if err != nil {
			log.Println("errrrrr", err)
			return c.Status(400).JSON(err)
		}
		return c.Status(200).JSON(response)

	})
}
