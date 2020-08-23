package businesses

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/api/businesses"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/test/helpers"
)

func TestGetProjects(t *testing.T) {
	log.Println("-- test get business --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/businesses"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	_, businessID := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should get the list of businesses for a certain user", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "businesses")
		assert.Nil(err)
		response := ress.(*businesses.ResponseList)
		assert.Len(response.Data, 1)
		assert.Equal(response.Total, 1)
	})

	t.Run("It should get the current business", func(t *testing.T) {
		assert := assert.New(t)
		signupData := helpers.FakeSignup[0]
		req := helpers.DoRequest("GET", path+"/current", nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "business")
		assert.Nil(err)
		response := ress.(*models.BusinessModel)
		assert.Equal(businessID, response.ID)
		assert.Equal(signupData.Domain, response.Domain)
		assert.Equal(signupData.Business, response.Name)

	})

}
