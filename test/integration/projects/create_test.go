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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateProject(t *testing.T) {
	log.Println("-- test create project --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/projects"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app, 1)

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
				ID:          primitive.NewObjectID(),
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
		log.Println("token", authResponse.Token)
		req := helpers.DoRequest("POST", path, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		response, err := helpers.GetResponseProject(res)
		assert.Nil(err)
		assert.Equal(authResponse.CurrentBusiness.ID, response.BusinessID)
		assert.Equal(authResponse.CurrentUser.User.ID, response.CreatedBy)
		assert.NotEmpty(response.ID)
		assert.Equal(fakeProject, response.Name)
		assert.EqualValues(fakeArea, response.Area)
		assert.Equal(enums.DefaultUnitArea, response.Unit)
		assert.Equal(fakeAddress.Country, response.Address.Country)
		assert.Equal(fakeAddress.State, response.Address.State)
		assert.Equal(fakeNotes, response.Notes)
		assert.Len(response.Images, 1)
		assert.Equal(fakeImages[0].ID, response.Images[0].ID)
		assert.Equal(fakeImages[0].Src, response.Images[0].Src)
		assert.Equal(fakeImages[0].Alt, response.Images[0].Alt)
		assert.Equal(fakeImages[0].Description, response.Images[0].Description)
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
		assert.EqualValues(0, response.Area)
		assert.Equal(enums.DefaultUnitArea, response.Unit)
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
			assert.Equal(services.T(errors.ErrInputInvalid), response.Message, response)
		}
	})

}