package errors

import (
	"log"

	"github.com/gofiber/fiber"
)

// ErrorHandler middleware
// This is a global handler for unexpected errors
func ErrorHandler() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		if c.Error() != nil {
			e := &HTTPError{}
			e.Code = 500
			e.Message = c.Error().Error()

			// TODO: log to external
			log.Println("-err", e)
			c.Status(e.Code).JSON(e)
		}
		c.Next()
	}
}
