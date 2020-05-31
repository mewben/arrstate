package pkg

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/pkg/auth"
)

// SetupBackend -
func SetupBackend(db *mongo.Database) *fiber.App {
	// db := startup.Init()

	app := fiber.New()

	// global middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Register routes
	auth.Routes(app, db)

	// Return the configured app
	return app
}
