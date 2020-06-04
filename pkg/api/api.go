package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/pkg/api/projects"
)

// Routes -
func Routes(app *fiber.App, db *mongo.Database) {
	g := app.Group("/api")

	// jwt middleware
	g.Use(jwtware.New(jwtware.Config{
		SigningKey:    []byte(viper.GetString("TOKEN_SIGNING_KEY")),
		SigningMethod: jwt.SigningMethodHS256.Name,
		ContextKey:    "user",
		TokenLookup:   "header:Authorization",
		AuthScheme:    "Bearer",
	}))

	projects.Routes(g, db)

}
