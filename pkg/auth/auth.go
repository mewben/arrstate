package auth

import (
	"context"
	"log"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/pkg/services/database"
)

// Handler -
type Handler struct {
	DB  *database.Service
	Ctx context.Context
}

// Routes -
func Routes(app *fiber.App, db *mongo.Database) {
	h := &Handler{
		DB: database.NewService(db),
	}

	app.Post("/auth/signup", func(c *fiber.Ctx) {
		h.Ctx = c.Fasthttp
		payload := &SignupPayload{}

		if err := c.BodyParser(&payload); err != nil {
			c.Next(err)
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
}
