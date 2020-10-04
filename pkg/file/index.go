package file

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/pkg/file/proxy"
)

// Routes -
func Routes(app *fiber.App, db *mongo.Database) {

	// TODO: permissions

	fileServer := viper.GetString("FILE_SERVER")

	log.Println("fileServer", fileServer)

	proxyConfig := proxy.Config{
		Hosts: fileServer,
		Before: func(c *fiber.Ctx) error {
			log.Println("beforee:", string(c.Request().RequestURI()))
			c.Request().SetRequestURI(strings.Replace(string(c.Request().RequestURI()), "/files", "", -1))
			log.Println("beforee2:", string(c.Request().RequestURI()))
			c.Request().Header.Set("X-Forwarded-Host", c.Hostname())
			c.Request().Header.Set("X-Forwarded-Proto", c.Protocol())

			return nil
		},
		// After: func(c *fiber.Ctx) error {
		// 	c.Set("X-Forwarded-Host", "localhost:8000")
		// 	return nil
		// },
	}

	// group := app.Group("/files")

	// group.Use(proxy.New(proxyConfig))

	app.Use("/files", proxy.New(proxyConfig))

	// app.Post("/files", jwtware.New(jwtware.Config{
	// 	SigningKey:    []byte(viper.GetString("TOKEN_SIGNING_KEY")),
	// 	SigningMethod: jwt.SigningMethodHS256.Name,
	// 	ContextKey:    "user",
	// 	TokenLookup:   "header:Authorization",
	// 	AuthScheme:    "Bearer",
	// }), proxy.New(proxyConfig))

	// app.Patch("/files", jwtware.New(jwtware.Config{
	// 	SigningKey:    []byte(viper.GetString("TOKEN_SIGNING_KEY")),
	// 	SigningMethod: jwt.SigningMethodHS256.Name,
	// 	ContextKey:    "user",
	// 	TokenLookup:   "header:Authorization",
	// 	AuthScheme:    "Bearer",
	// }), proxy.New(proxyConfig))

}
