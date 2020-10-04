package pkg

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/pkg/api"
	"github.com/mewben/arrstate/pkg/auth"
	"github.com/mewben/arrstate/pkg/file"
)

// SetupBackend -
func SetupBackend(db *mongo.Database) *fiber.App {
	app := fiber.New()
	// app := fiber.New(fiber.Config{
	// 	// Override default error handler
	//   ErrorHandler: func(ctx *fiber.Ctx, err error) error {
	//       // Statuscode defaults to 500
	//       code := fiber.StatusInternalServerError

	//       // Retreive the custom statuscode if it's an fiber.*Error
	//       if e, ok := err.(*fiber.Error); ok {
	//           code = e.Code
	//       }

	//       // Send custom error page
	//       err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
	//       if err != nil {
	//           // In case the SendFile fails
	//           return ctx.Status(500).SendString("Internal Server Error")
	//       }

	//       // Return from handler
	//       return nil
	//   }
	// })

	// global middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Register routes
	auth.Routes(app, db)
	api.Routes(app, db)

	// file-uploads proxy
	file.Routes(app, db)

	// static
	app.Static("/", "./web/public")
	app.Get("/*", func(c *fiber.Ctx) error {
		if strings.Contains(c.Path(), "/api") {
			return c.Next()
		}

		if err := c.SendFile("./web/public/index.html"); err != nil {
			return fiber.ErrInternalServerError
		}

		return nil
	})

	// app.Use(errors.ErrorHandler())

	// Return the configured app
	return app
}
