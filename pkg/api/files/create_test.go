package files_test

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/test/helpers"
)

func TestCreateFile(t *testing.T) {
	log.Println("-- test file.post")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/files"

	// setup
	helpers.CleanupFixture(db)
	token := helpers.SignupFixture(app, 0)
	userID, businessID, _ := helpers.CheckJWT(token, assert.New(t))

	t.Run("It should create file data", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"url":        "http://url.com",
			"title":      "Title",
			"ext":        "jpg",
			"size":       12343,
			"mimeType":   "image/*",
			"entityType": "invoice",
			"entityID":   primitive.NewObjectID().Hex(),
			"type":       "file",
			"link":       "",
		}

		req := helpers.DoRequest("POST", path, data, token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		response, err := GetResponse(res)
		assert.Nil(err)
		assert.NotZero(response.ID)
		assert.Equal(userID, response.CreatedBy)
		assert.Equal(businessID, response.BusinessID)

		assert.Equal(data["url"], response.URL)
		assert.Equal(data["title"], response.Title)
		assert.Equal(data["ext"], response.Extension)
		assert.EqualValues(data["size"], response.Size)
		assert.Equal(data["ext"], response.Extension)
		assert.Equal(data["entityType"], *response.EntityType)
		assert.Equal(data["entityID"], response.EntityID.Hex())
		assert.Equal(data["type"], response.Type)
		assert.Equal(data["link"], response.Link)

	})
}
