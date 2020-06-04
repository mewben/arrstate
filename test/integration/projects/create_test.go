package projects

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
)

func TestCreateProject(t *testing.T) {
	log.Println("-- test create project --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/projects"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app)

	t.Run("It should create project", func(t *testing.T) {
		assert := assert.New(t)
		fakeProject := "testproj"
		fakeAddress := models.NewAddressModel()
		fakeAddress.Country = "PH"
		fakeAddress.State = "Bohol"
		fakeArea := 36.5
		fakeNotes := "Sample Notes"
		fakeImages := []*models.ImageModel{
			{
				ID:          "id",
				Src:         "src",
				Alt:         "alt",
				Description: "description",
			},
		}
		data := fiber.Map{
			"name":    fakeProject,
			"address": fakeAddress,
			"area":    fakeArea,
			"unit":    enums.DefaultUnitArea,
			"notes":   fakeNotes,
			"images":  fakeImages,
		}
		req := helpers.DoRequest("POST", path, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		response, err := helpers.GetResponseProject(res)
		assert.Nil(err)
		assert.Equal(response.BusinessID, authResponse.CurrentBusiness.ID)
		assert.Equal(response.CreatedBy, authResponse.CurrentUser.User.ID)
		assert.NotEmpty(response.ID)
		assert.Equal(response.Name, fakeProject)
		assert.EqualValues(response.Area, fakeArea)
		assert.Equal(response.Unit, enums.DefaultUnitArea)
		assert.Equal(response.Address.Country, fakeAddress.Country)
		assert.Equal(response.Address.State, fakeAddress.State)
		assert.Equal(response.Notes, fakeNotes)
		assert.Len(response.Images, 1)
		assert.Equal(response.Images[0].ID, fakeImages[0].ID)
		assert.Equal(response.Images[0].Src, fakeImages[0].Src)
		assert.Equal(response.Images[0].Alt, fakeImages[0].Alt)
		assert.Equal(response.Images[0].Description, fakeImages[0].Description)
	})

	t.Run("It should set defaultValues if not set", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"name": "Project Name 2",
		}
		req := helpers.DoRequest("POST", path, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		response, err := helpers.GetResponseProject(res)
		assert.Nil(err)
		assert.EqualValues(response.Area, 0)
		assert.Equal(response.Unit, enums.DefaultUnitArea)
	})

	t.Run("It should validate project inputs", func(t *testing.T) {
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
			req := helpers.DoRequest("POST", path, payload, authResponse.Token)
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, payload)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(response.Message, services.T(errors.ErrInputInvalid), response)
		}
	})

}
