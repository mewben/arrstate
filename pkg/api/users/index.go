package users

import (
	"context"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

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
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required,min=6"`
	DeviceCode string `json:"-"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(g fiber.Router, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Get("/users/current", func(c *fiber.Ctx) error {
		var err error
		h.Ctx = c.Context()
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			return c.Status(400).JSON(err)
		}

		return c.Status(200).JSON(h.User)
	})
}
