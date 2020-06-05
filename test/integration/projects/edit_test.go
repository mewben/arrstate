package projects

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestEditProject(t *testing.T) {
	log.Println("-- test edit project --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/projects"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app, 1)
	_, authResponse2 := helpers.SignupFixture(app, 2)
	project := helpers.ProjectFixture(app, authResponse.Token, 1)
	project2 := helpers.ProjectFixture(app, authResponse2.Token, 2)

	t.Run("It should edit project", func(t *testing.T) {
		assert := assert.New(t)
		updName := "edit1"
		updAddress := models.NewAddressModel()
		updAddress.Country = "US"
		updAddress.State = "Ohio"
		updArea := 24
		updUnit := "sq.in"
		updNotes := "Edit notes"
		data := fiber.Map{
			"name":    updName,
			"address": updAddress,
			"area":    updArea,
			"unit":    updUnit,
			"notes":   updNotes,
		}
		req := helpers.DoRequest("PUT", path+"/"+project.ID.Hex(), data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseProject(res)
		assert.Nil(err)
		assert.Equal(authResponse.CurrentBusiness.ID, response.BusinessID)
		assert.Equal(authResponse.CurrentUser.User.ID, response.CreatedBy)
		assert.Equal(project.ID, response.ID)
		assert.Equal(updName, response.Name)
		assert.EqualValues(updArea, response.Area)
		assert.Equal(updUnit, response.Unit)
		assert.Equal(updAddress.Country, response.Address.Country)
		assert.Equal(updAddress.State, response.Address.State)
		assert.Equal(updNotes, response.Notes)
		assert.Equal(authResponse.CurrentUser.User.ID, response.UpdatedBy)

	})

	t.Run("It should validate input", func(t *testing.T) {
		assert := assert.New(t)
		payloads := []fiber.Map{
			{},
			{
				"name": "",
			},
			{
				"name": "testp",
				"area": "notanumber",
			},
			{
				"name": "testp",
				"area": "100",
			},
			{
				"name": "testp",
				"area": -102,
			},
		}

		for _, payload := range payloads {
			req := helpers.DoRequest("PUT", path+"/"+project.ID.Hex(), payload, authResponse.Token)
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, payload)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrInputInvalid), response.Message, response)
		}

	})

	t.Run("Permissions", func(t *testing.T) {

		t.Run("It should not edit if not in the current business", func(t *testing.T) {
			assert := assert.New(t)
			data := fiber.Map{
				"name": "edit another",
			}
			req := helpers.DoRequest("PUT", path+"/"+project2.ID.Hex(), data, authResponse.Token)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)

		})
	})
}
