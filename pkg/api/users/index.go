package users

import (
	"context"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"

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
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required,min=6"`
	DeviceCode string `json:"-"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(g *fiber.Group, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Get("/me", func(c *fiber.Ctx) {
		var err error
		h.Ctx = c.Fasthttp
		h.User, h.Business, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.GetMe()
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)
	})
}
