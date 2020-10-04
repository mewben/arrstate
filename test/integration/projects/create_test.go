package projects

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestCreateProject(t *testing.T) {
	log.Println("-- test create project --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/projects"

	// setup
	helpers.CleanupFixture(db)
	token := helpers.SignupFixture(app, 0)
	userID, businessID, _ := helpers.CheckJWT(token, assert.New(t))

	t.Run("It should create project", func(t *testing.T) {
		assert := assert.New(t)
		fakeProject := "testproj"
		fakeAddress := models.NewAddressModel()
		fakeAddress.Country = "PH"
		fakeAddress.State = "Bohol"
		fakeArea := 36.5
		fakeNotes := "Sample Notes"
		// fakeImages := []*models.FileSchemaWID{
		// 	{
		// 		ID:        primitive.NewObjectID(),
		// 		URL:       "url",
		// 		Title:     "alt",
		// 		Extension: "jpg",
		// 		Size:      1235,
		// 		MimeType:  "image/*",
		// 	},
		// }
		data := fiber.Map{
			"name":    fakeProject,
			"address": fakeAddress,
			"area":    fakeArea,
			"unit":    enums.DefaultUnitArea,
			"notes":   fakeNotes,
			// "files":   fakeImages,
		}
		req := helpers.DoRequest("POST", path, data, token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "project")
		assert.Nil(err)
		response := ress.(*models.ProjectModel)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.CreatedBy)
		assert.Equal(userID, response.UpdatedBy)
		assert.False(response.ID.IsZero())
		assert.Equal(fakeProject, response.Name)
		assert.EqualValues(fakeArea, response.Area)
		assert.Equal(enums.DefaultUnitArea, response.Unit)
		assert.Equal(fakeAddress.Country, response.Address.Country)
		assert.Equal(fakeAddress.State, response.Address.State)
		assert.Equal(fakeNotes, response.Notes)
		// assert.Len(response.Files, 1)
		// assert.Equal(fakeImages[0].ID, response.Files[0].ID)
		// assert.Equal(fakeImages[0].URL, response.Files[0].URL)
		// assert.Equal(fakeImages[0].Title, response.Files[0].Title)
		// assert.Equal(fakeImages[0].Extension, response.Files[0].Extension)
		// assert.Equal(fakeImages[0].MimeType, response.Files[0].MimeType)
		// assert.Equal(fakeImages[0].Size, response.Files[0].Size)
	})

	t.Run("It should set defaultValues if not set", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"name": "Project Name 2",
		}
		req := helpers.DoRequest("POST", path, data, token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "project")
		assert.Nil(err)
		response := ress.(*models.ProjectModel)
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
			req := helpers.DoRequest("POST", path, payload, token)
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
