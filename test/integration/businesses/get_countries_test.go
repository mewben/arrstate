package businesses

import (
	"log"
	"os"
	"testing"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestGetCountries(t *testing.T) {
	log.Println("-- BUSINESS.GET.COUNTRIES --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/businesses/countries"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)

	t.Run("It should get countries", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path, nil, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		assert.Len(response["countries"], 249)
	})
}
