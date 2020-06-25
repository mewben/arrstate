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
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	project := helpers.ProjectFixture(app, token1, 0)
	project2 := helpers.ProjectFixture(app, token2, 1)
	lot1 := helpers.LotFixture(app, token1, &project.ID, 0)
	lot2 := helpers.LotFixture(app, token2, &project2.ID, 1)

	t.Run("It should remove lot", func(t *testing.T) {
		lotID := lot1.ID.Hex()
		assert := assert.New(t)
		req := helpers.DoRequest("DELETE", path+"/"+lotID, nil, token1)

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
			req := helpers.DoRequest("DELETE", path+"/"+lotID, nil, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)
		})
	})
}
