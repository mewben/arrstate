package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ExtractClaims -
func ExtractClaims(c *fiber.Ctx) (userID, businessID, personID primitive.ObjectID, err error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, err = primitive.ObjectIDFromHex(claims["sub"].(string))
	if err != nil {
		return
	}
	businessID, err = primitive.ObjectIDFromHex(claims["businessID"].(string))
	if err != nil {
		return
	}

	personID, err = primitive.ObjectIDFromHex(claims["personID"].(string))
	return
}
