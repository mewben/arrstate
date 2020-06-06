package lots

import (
	"context"
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	models.LotModel
}

// Payload2 -
type Payload2 struct {
	ProjectID  primitive.ObjectID `json:"projectID"`
	Name       string             `json:"name" validate:"required"`
	Area       float32            `json:"area" validate:"number,min=0"`
	Price      float32            `json:"price" validate:"number,min=0"`
	PriceAddon float32            `json:"priceAddon" validate:"number,min=0"`
	models.MetaModel
}

// ResponseList -
type ResponseList struct {
	Total int                `json:"total"`
	Data  []*models.LotModel `json:"list"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()

// Routes -
func Routes(g *fiber.Group, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	// g.Get("/lots", func(c *fiber.Ctx) {
	// 	log.Println("lots.get")
	// 	if err := h.Prepare(c); err != nil {
	// 		c.Status(400).JSON(err)
	// 		return
	// 	}

	// 	response, err := h.Get()
	// 	if err != nil {
	// 		log.Println("errrrrr", err)
	// 		c.Status(400).JSON(err)
	// 		return
	// 	}
	// 	c.Status(200).JSON(response)
	// })

	g.Post("/lots", func(c *fiber.Ctx) {
		log.Println("lots.post")
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
		log.Println("aferecreate")
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(201).JSON(response)
	})

	g.Put("/lots", func(c *fiber.Ctx) {
		log.Println("lots.put")
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

		response, err := h.Edit(payload)
		if err != nil {
			log.Println("errrrrr", err)
			c.Status(400).JSON(err)
			return
		}
		c.Status(200).JSON(response)

	})

	// g.Delete("/lots/:projectID", func(c *fiber.Ctx) {
	// 	log.Println("lots.delete")
	// 	if err := h.Prepare(c); err != nil {
	// 		c.Status(400).JSON(err)
	// 		return
	// 	}

	// 	response, err := h.Remove(c.Params("projectID"))
	// 	if err != nil {
	// 		log.Println("errrrrr", err)
	// 		c.Status(400).JSON(err)
	// 		return
	// 	}
	// 	c.Status(200).JSON(response)

	// })
}
