package errors

// import (
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// )

// // ErrorHandler middleware
// // This is a global handler for unexpected errors
// func ErrorHandler() func(*fiber.Ctx) {
// 	return func(c *fiber.Ctx) error {
// 		if c.Error() != nil {
// 			log.Println("-- Error Handler: ", c.Error())
// 			e := &HTTPError{}
// 			e.Code = 500
// 			e.Message = c.Error().Error()

// 			// TODO: log to external
// 			return c.Status(e.Code).JSON(e)
// 		}
// 		c.Next()
// 	}
// }
