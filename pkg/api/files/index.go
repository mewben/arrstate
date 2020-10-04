package files

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
	DB       *database.Service
	Ctx      context.Context
	User     *models.UserModel
	Business *models.BusinessModel
}

// Params -
type Params struct {
	EntityType *string              `query:"entityType"`
	EntityID   *models.NullObjectID `query:"entityID"`
}

// Payload -
type Payload struct {
	models.FileModel
}

// ResponseList -
type ResponseList struct {
	Total int                 `json:"total"`
	Data  []*models.FileModel `json:"list"`
}

var validate = validator.New()

// Routes -
func Routes(g fiber.Router, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Get("/files", func(c *fiber.Ctx) error {
		var err error
		h.Ctx = c.Context()
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			return c.Status(400).JSON(err)
		}

		params := &Params{}
		if err := c.QueryParser(params); err != nil {
			log.Println("errrbodyparser get params:", err)
			return c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
		}

		response, err := h.Get(params)
		if err != nil {
			log.Println("errrrrr", err)
			return c.Status(400).JSON(err)
		}
		return c.Status(200).JSON(response)
	})

	g.Post("/files", func(c *fiber.Ctx) error {
		log.Println("files.post")
		var err error
		h.Ctx = c.Context()
		h.User, h.Business, _, err = utils.PrepareHandler(c, h.DB)
		if err != nil {
			return c.Status(400).JSON(err)
		}

		payload := &Payload{}
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
