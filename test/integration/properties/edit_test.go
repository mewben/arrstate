package properties

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestEditProperty(t *testing.T) {
	log.Println("-- test edit property --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/properties"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	project := helpers.ProjectFixture(app, token1, 0)
	project2 := helpers.ProjectFixture(app, token2, 1)
	property1 := helpers.PropertyFixture(app, token1, &project.ID, 0)
	property2 := helpers.PropertyFixture(app, token2, &project2.ID, 1)
	userID, businessID := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should edit property", func(t *testing.T) {
		assert := assert.New(t)
		updName := "edit"
		updType := enums.PropertyTypeHouse
		updArea := 12.5
		updPrice := 100.5
		updPriceAddon := 15.3
		updNotes := "edit notes"
		data := fiber.Map{
			"_id":        property1.ID,
			"name":       updName,
			"type":       updType,
			"area":       updArea,
			"price":      updPrice,
			"priceAddon": updPriceAddon,
			"notes":      updNotes,
		}
		req := helpers.DoRequest("PUT", path, data, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "property")
		assert.Nil(err)
		response := ress.(*models.PropertyModel)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.CreatedBy)
		assert.Equal(property1.ID, response.ID)
		assert.Equal(project.ID, *response.ProjectID)
		assert.Equal(updName, response.Name)
		assert.Equal(updType, response.Type)
		assert.EqualValues(updArea, response.Area)
		assert.EqualValues(updPrice, response.Price)
		assert.EqualValues(updPriceAddon, response.PriceAddon)
		assert.Equal(updNotes, response.Notes)
		assert.Equal(userID, response.UpdatedBy)

	})

	t.Run("It should validate input", func(t *testing.T) {
		assert := assert.New(t)
		payloads := []fiber.Map{
			{},
			{
				"_id": "",
			},
			{
				"_id":  "invalidid",
				"name": "testp",
			},
			{
				"_id":  property1.ID,
				"name": "",
			},
			{
				"_id":  property1.ID,
				"name": "testp",
				"area": "notanumber",
			},
			{
				"_id":  property1.ID,
				"name": "testp",
				"area": "100",
			},
			{
				"_id":  property1.ID,
				"name": "testp",
				"area": -102,
			},
			{
				"_id":   property1.ID,
				"name":  "testp",
				"area":  102,
				"price": -14,
			},
			{
				"_id":        property1.ID,
				"name":       "testp",
				"area":       102,
				"price":      14,
				"priceAddon": -14,
			},
			{
				"_id":        property1.ID,
				"name":       "testp",
				"type":       "",
				"area":       102,
				"price":      14,
				"priceAddon": 14,
			},
			{
				"_id":        property1.ID,
				"name":       "testp",
				"type":       "invalidtype",
				"area":       102,
				"price":      14,
				"priceAddon": 14,
			},
		}

		for _, payload := range payloads {
			req := helpers.DoRequest("PUT", path, payload, token1)
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, payload)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrInputInvalid), response.Message, response)
		}

	})

	t.Run("It should not edit projectID", func(t *testing.T) {
		assert := assert.New(t)
		updName := "edit2"
		data := fiber.Map{
			"_id":       property1.ID,
			"name":      updName,
			"type":      enums.PropertyTypeHouse,
			"projectID": primitive.NewObjectID(),
		}
		req := helpers.DoRequest("PUT", path, data, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "property")
		assert.Nil(err)
		response := ress.(*models.PropertyModel)
		assert.Equal(property1.ID, response.ID)
		assert.Equal(project.ID, *response.ProjectID)
		assert.Equal(updName, response.Name)
	})

	t.Run("Permissions", func(t *testing.T) {

		t.Run("It should not edit from other business", func(t *testing.T) {
			assert := assert.New(t)
			data := fiber.Map{
				"_id":  property2.ID,
				"name": "edit another",
				"type": enums.PropertyTypeHouse,
			}
			req := helpers.DoRequest("PUT", path, data, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)

		})
	})

}
