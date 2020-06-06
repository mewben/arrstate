package lots

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
)

func TestRemoveLot(t *testing.T) {
	log.Println("-- test remove lot --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/lots"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app, 1)
	_, authResponse2 := helpers.SignupFixture(app, 2)
	project := helpers.ProjectFixture(app, authResponse.Token, 1)
	project2 := helpers.ProjectFixture(app, authResponse2.Token, 2)
	lot1 := helpers.LotFixture(app, authResponse.Token, project.ID, 1)
	lot2 := helpers.LotFixture(app, authResponse2.Token, project2.ID, 2)

	t.Run("It should remove lot", func(t *testing.T) {
		lotID := lot1.ID.Hex()
		assert := assert.New(t)
		req := helpers.DoRequest("DELETE", path+"/"+lotID, nil, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseDelete(res)
		assert.Nil(err)
		assert.Equal(lotID, response["lot"])

	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not remove lot from other business", func(t *testing.T) {
			lotID := lot2.ID.Hex()
			assert := assert.New(t)
			req := helpers.DoRequest("DELETE", path+"/"+lotID, nil, authResponse.Token)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)
		})
	})
}
