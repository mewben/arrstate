package people

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestEditPersonLocale(t *testing.T) {
	log.Println("-- test edit person locale --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/people"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	person1 := helpers.PersonFixture(app, token1, 0) // person1
	person2 := helpers.PersonFixture(app, token2, 1)
	_, _, personID := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should edit my locale", func(t *testing.T) {
		assert := assert.New(t)
		data := map[string]string{
			"language":        "fr",
			"dateFormat":      "YYYY-MM-DD",
			"timeFormat":      "hh:mm",
			"timestampFormat": "relative",
		}
		req := helpers.DoRequest("PUT", path+"/"+personID.Hex()+"/locale", data, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "person")
		assert.Nil(err)
		response := ress.(*models.PersonModel)
		assert.Equal(personID, response.ID)

		locale := response.Locale
		assert.Equal(data["language"], locale.Language)
		assert.Equal(data["dateFormat"], locale.DateFormat)
		assert.Equal(data["timeFormat"], locale.TimeFormat)
		assert.Equal(data["timestampFormat"], locale.TimestampFormat)
	})

	t.Run("It should edit not edit locale of another person", func(t *testing.T) {
		assert := assert.New(t)
		data := map[string]string{
			"language":        "fr",
			"dateFormat":      "YYYY-MM-DD",
			"timeFormat":      "hh:mm",
			"timestampFormat": "relative",
		}
		cases := []struct {
			personID string
		}{
			{
				person1.ID.Hex(),
			},
			{
				person2.ID.Hex(),
			},
		}

		for _, item := range cases {
			req := helpers.DoRequest("PUT", path+"/"+item.personID+"/locale", data, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)
		}
	})
}
