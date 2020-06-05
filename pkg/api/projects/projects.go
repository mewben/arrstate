package projects

import (
	"context"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
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
	Name    string              `json:"name" validate:"required"`
	Address models.AddressModel `json:"address"`
	Area    float32             `json:"area" validate:"number,min=0"`
	Unit    string              `json:"unit"`
	models.MetaModel
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(g *fiber.Group, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	g.Post("/projects", func(c *fiber.Ctx) {
		log.Println("projects.post")
		if err := h.Prepare(c); err != nil {
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

	g.Put("/projects/:projectID", func(c *fiber.Ctx) {
		log.Println("projects.put")
		if err := h.Prepare(c); err != nil {
			c.Status(400).JSON(err)
			return
		}

		payload := &Payload{}
		if err := c.BodyParser(&payload); err != nil {
			log.Println("errrbodyparser", err)
			c.Status(400).JSON(errors.NewHTTPError(errors.ErrInputInvalid, err))
			return
		}

		response, err := h.Edit(c.Params("projectID"), payload)
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})

	g.Delete("/projects/:projectID", func(c *fiber.Ctx) {
		log.Println("projects.delete")
		if err := h.Prepare(c); err != nil {
			c.Status(400).JSON(err)
			return
		}

		response, err := h.Remove(c.Params("projectID"))
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})
}

// Prepare -
func (h *Handler) Prepare(c *fiber.Ctx) error {
	h.Ctx = c.Fasthttp
	userID, businessID, err := utils.ExtractClaims(c)
	if err != nil {
		return err
	}

	// get user
	h.User = h.DB.FindUser(h.Ctx, userID)
	if h.User == nil {
		return errors.NewHTTPError(errors.ErrNotFoundUser)
	}

	// get business
	h.Business = h.DB.FindBusiness(h.Ctx, businessID)
	if h.Business == nil {
		return errors.NewHTTPError(errors.ErrNotFoundBusiness)
	}

	return nil
}
