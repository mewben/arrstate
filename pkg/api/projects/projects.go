package projects

import (
	"context"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/internal/enums"
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
}

// Prepare -
func (h *Handler) Prepare(c *fiber.Ctx) error {
	h.Ctx = c.Fasthttp
	userID, businessID := utils.ExtractClaims(c)

	// get user
	foundUser := h.DB.FindByID(h.Ctx, enums.CollUsers, userID)
	if foundUser == nil {
		return errors.NewHTTPError(errors.ErrNotFoundUser)
	}
	h.User = foundUser.(*models.UserModel)

	// get business
	foundBusiness := h.DB.FindByID(h.Ctx, enums.CollBusinesses, businessID)
	if foundBusiness == nil {
		return errors.NewHTTPError(errors.ErrNotFoundBusiness)
	}
	h.Business = foundBusiness.(*models.BusinessModel)

	return nil
}
