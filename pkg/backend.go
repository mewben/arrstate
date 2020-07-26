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

func init() {
	// include in pkger

}

// SetupBackend -
func SetupBackend(db *mongo.Database) *fiber.App {
	// db := startup.Init()

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

	app.Use(errors.ErrorHandler())

	// Return the configured app
	return app
}
