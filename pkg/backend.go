package pkg

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/pkg/api"
	"github.com/mewben/arrstate/pkg/auth"
	"github.com/mewben/arrstate/pkg/errors"
)

// SetupBackend -
func SetupBackend(db *mongo.Database) *fiber.App {
	app := fiber.New()

	// global middleware
	app.Use(logger.New())
	app.Use(recover.New())
	// app.Use(cors.New())

	// Register routes
	auth.Routes(app, db)
	api.Routes(app, db)

	// static
	app.Static("/", "./web/public")
	app.Get("/*", func(c *fiber.Ctx) {
		if err := c.SendFile("./web/public/index.html"); err != nil {
			c.Next(fiber.ErrInternalServerError)
		}
	})

	app.Use(errors.ErrorHandler())

	// Return the configured app
	return app
}
