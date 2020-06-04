package projects

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/models"
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
	_, authResponse := helpers.SignupFixture(app)
	project := helpers.ProjectFixture(app, authResponse.Token)

	log.Println("project", project.ID)

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
		log.Println("token", authResponse.Token)
		req := helpers.DoRequest("PUT", path+"/"+project.ID, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(203, res.StatusCode, res)

	})

	t.Run("It should edit only those set", func(t *testing.T) {

	})

	t.Run("It should validate input", func(t *testing.T) {

	})
}
