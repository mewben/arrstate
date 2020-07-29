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

func TestGetInitValues(t *testing.T) {
	log.Println("-- BUSINESS.GET.INIT_VALUES --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/businesses"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)

	t.Run("It should get countries", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path+"/countries", nil, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		assert.Len(response["countries"], 249)
	})

	t.Run("It should get currencies", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path+"/currencies", nil, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		assert.Len(response["currencies"], 116)
	})
}
