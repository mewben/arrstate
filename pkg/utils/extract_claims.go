package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

// ExtractClaims -
func ExtractClaims(c *fiber.Ctx) (userID, businessID string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID = claims["sub"].(string)
	businessID = claims["businessID"].(string)
	return
}
