package properties

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestCreateProperty(t *testing.T) {
	log.Println("-- test create property --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/properties"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	project := helpers.ProjectFixture(app, token1, 0)
	userID, businessID, _ := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should create property", func(t *testing.T) {
		assert := assert.New(t)
		fakeProjectID := project.ID.Hex()
		fakeName := "Fake property"
		fakeType := enums.PropertyTypeLot
		fakeArea := 53.5
		fakePrice := 10050
		fakePriceAddon := 10100
		fakeNotes := "Sample Notes"
		fakeImages := []*models.FileSchemaWID{
			{
				ID:        primitive.NewObjectID(),
				URL:       "url",
				Title:     "alt",
				Extension: "jpg",
				Size:      1235,
				MimeType:  "image/*",
			},
		}

		data := fiber.Map{
			"projectID":  fakeProjectID,
			"name":       fakeName,
			"type":       fakeType,
			"area":       fakeArea,
			"price":      fakePrice,
			"priceAddon": fakePriceAddon,
			"notes":      fakeNotes,
			"files":      fakeImages,
		}
		req := helpers.DoRequest("POST", path, data, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "property")
		assert.Nil(err)
		response := ress.(*models.PropertyModel)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.CreatedBy)
		assert.False(response.ID.IsZero())
		assert.Equal(fakeProjectID, response.ProjectID.Hex())
		assert.Equal(fakeName, response.Name)
		assert.Equal(fakeType, response.Type)
		assert.Equal(enums.StatusAvailable, response.Status)
		assert.EqualValues(fakeArea, response.Area)
		assert.EqualValues(fakePrice, response.Price)
		assert.EqualValues(fakePriceAddon, response.PriceAddon)
		assert.Equal(fakeNotes, response.Notes)
		assert.Len(response.Files, 1)
		assert.Equal(fakeImages[0].ID, response.Files[0].ID)
		assert.Equal(fakeImages[0].URL, response.Files[0].URL)
		assert.Equal(fakeImages[0].Title, response.Files[0].Title)
		assert.Equal(fakeImages[0].Extension, response.Files[0].Extension)
		assert.Equal(fakeImages[0].MimeType, response.Files[0].MimeType)
		assert.Equal(fakeImages[0].Size, response.Files[0].Size)
	})

	t.Run("It should create property without projectID", func(t *testing.T) {
		assert := assert.New(t)
		fakeName := "Fake property2"
		fakeArea := 5350
		fakePrice := 10050
		fakePriceAddon := 1010
		fakeNotes := "Sample Notes"
		fakeImages := []*models.FileSchemaWID{
			{
				ID:        primitive.NewObjectID(),
				URL:       "url",
				Title:     "alt",
				Extension: "jpg",
				Size:      1235,
				MimeType:  "image/*",
			},
		}

		data := fiber.Map{
			"name":       fakeName,
			"type":       enums.PropertyTypeHouse,
			"area":       fakeArea,
			"price":      fakePrice,
			"priceAddon": fakePriceAddon,
			"notes":      fakeNotes,
			"files":      fakeImages,
		}
		req := helpers.DoRequest("POST", path, data, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "property")
		assert.Nil(err)
		response := ress.(*models.PropertyModel)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.CreatedBy)
		assert.False(response.ID.IsZero())
		assert.Nil(response.ProjectID)
		assert.Equal(fakeName, response.Name)
		assert.Equal(enums.PropertyTypeHouse, response.Type)
		assert.EqualValues(fakeArea, response.Area)
		assert.EqualValues(fakePrice, response.Price)
		assert.EqualValues(fakePriceAddon, response.PriceAddon)
		assert.Equal(fakeNotes, response.Notes)
		assert.Len(response.Files, 1)
		assert.Equal(fakeImages[0].ID, response.Files[0].ID)
		assert.Equal(fakeImages[0].URL, response.Files[0].URL)
		assert.Equal(fakeImages[0].Title, response.Files[0].Title)
		assert.Equal(fakeImages[0].Extension, response.Files[0].Extension)
		assert.Equal(fakeImages[0].MimeType, response.Files[0].MimeType)
		assert.Equal(fakeImages[0].Size, response.Files[0].Size)
	})

	t.Run("It should validate property inputs", func(t *testing.T) {
		assert := assert.New(t)
		projectID := project.ID.Hex()
		cases := []struct {
			err     string
			payload fiber.Map
		}{
			{
				errors.ErrInputInvalid,
				fiber.Map{},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"name": "",
					"i":    2,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"projectID": "",
					"name":      "testproperty",
					"i":         3,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"projectID": "invalidid",
					"name":      "testproperty",
					"i":         4,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"projectID": projectID,
					"name":      "testproperty",
					"area":      "notanumber",
					"i":         5,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"projectID": projectID,
					"name":      "testproperty",
					"area":      "100",
					"i":         6,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"projectID": projectID,
					"name":      "testproperty",
					"area":      -100,
					"i":         7,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"projectID": projectID,
					"name":      "testproperty",
					"area":      100.5,
					"price":     -100,
					"i":         8,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"projectID":  projectID,
					"name":       "testproperty",
					"area":       100.5,
					"price":      10090,
					"priceAddon": -40.5,
					"i":          9,
				},
			},
			{
				errors.ErrNotFoundProject,
				fiber.Map{
					"projectID":  primitive.NewObjectID(),
					"name":       "testproperty",
					"type":       enums.PropertyTypeHouse,
					"area":       100.5,
					"price":      10090,
					"priceAddon": 4050,
					"i":          10,
				},
			},
			{
				errors.ErrInputInvalid, // no property type
				fiber.Map{
					"projectID":  projectID,
					"name":       "testproperty",
					"type":       "",
					"area":       100.5,
					"price":      10090,
					"priceAddon": 4050,
					"i":          11,
				},
			},
			{
				errors.ErrInputInvalid, // invalid property type
				fiber.Map{
					"projectID":  projectID,
					"name":       "testproperty",
					"type":       "someunknowntype",
					"area":       100.5,
					"price":      10090,
					"priceAddon": 4050,
					"i":          12,
				},
			},
		}

		for _, item := range cases {
			req := helpers.DoRequest("POST", path, item.payload, token1)
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, item.payload)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(item.err), response.Message, item.payload)
		}
	})

}
