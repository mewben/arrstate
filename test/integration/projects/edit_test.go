package projects

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
)

func TestEditProject(t *testing.T) {
	log.Println("-- test edit project --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/projects"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	project := helpers.ProjectFixture(app, token1, 0)
	project2 := helpers.ProjectFixture(app, token2, 1)
	userID, businessID := helpers.CheckJWT(token1, assert.New(t))

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
			"_id":     project.ID,
			"name":    updName,
			"address": updAddress,
			"area":    updArea,
			"unit":    updUnit,
			"notes":   updNotes,
		}
		req := helpers.DoRequest("PUT", path, data, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "project")
		assert.Nil(err)
		response := ress.(*models.ProjectModel)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.CreatedBy)
		assert.Equal(project.ID, response.ID)
		assert.Equal(updName, response.Name)
		assert.EqualValues(updArea, response.Area)
		assert.Equal(updUnit, response.Unit)
		assert.Equal(updAddress.Country, response.Address.Country)
		assert.Equal(updAddress.State, response.Address.State)
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
				"_id":  project.ID,
				"name": "",
			},
			{
				"_id":  project.ID,
				"name": "testp",
				"area": "notanumber",
			},
			{
				"_id":  project.ID,
				"name": "testp",
				"area": "100",
			},
			{
				"_id":  project.ID,
				"name": "testp",
				"area": -102,
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

	t.Run("Permissions", func(t *testing.T) {

		t.Run("It should not edit if not in the current business", func(t *testing.T) {
			assert := assert.New(t)
			data := fiber.Map{
				"_id":  project2.ID,
				"name": "edit another",
			}
			req := helpers.DoRequest("PUT", path, data, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFoundProject), response.Message, response)

		})
	})
}
