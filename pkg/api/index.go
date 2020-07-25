package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/pkg/api/blocks"
	"github.com/mewben/arrstate/pkg/api/businesses"
	"github.com/mewben/arrstate/pkg/api/invoices"
	"github.com/mewben/arrstate/pkg/api/people"
	"github.com/mewben/arrstate/pkg/api/projects"
	"github.com/mewben/arrstate/pkg/api/properties"
	"github.com/mewben/arrstate/pkg/api/users"
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

	businesses.Routes(g, db)
	users.Routes(g, db)
	projects.Routes(g, db)
	properties.Routes(g, db)
	people.Routes(g, db)
	invoices.Routes(g, db)
	blocks.Routes(g, db)

}
