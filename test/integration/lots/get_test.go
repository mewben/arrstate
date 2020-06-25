package lots

import (
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/api/lots"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
)

func TestGetLots(t *testing.T) {
	log.Println("-- test get lots --")

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
	helpers.LotFixture(app, token1, &project.ID, 1)
	helpers.LotFixture(app, token1, nil, 1)
	lot4 := helpers.LotFixture(app, token2, &project2.ID, 1)

	t.Run("It should get the list of lots", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "lots")
		assert.Nil(err)
		response := ress.(*lots.ResponseList)
		assert.Len(response.Data, 3)
		assert.Equal(response.Total, 3)
	})

	t.Run("It should get the list of lots with projectID", func(t *testing.T) {
		assert := assert.New(t)
		query := url.Values{}
		query.Add("projectID", project.ID.Hex())
		req := helpers.DoRequest("GET", path+"?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "lots")
		assert.Nil(err)
		response := ress.(*lots.ResponseList)
		assert.Len(response.Data, 2)
		assert.Equal(response.Total, 2)
	})

	t.Run("It should get the list of lots with projectID = null", func(t *testing.T) {
		assert := assert.New(t)
		query := url.Values{}
		query.Add("projectID", "null")
		req := helpers.DoRequest("GET", path+"?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "lots")
		assert.Nil(err)
		response := ress.(*lots.ResponseList)
		assert.Len(response.Data, 1)
		assert.Equal(response.Total, 1)
	})

	t.Run("It should get the lot by id", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path+"/"+lot1.ID.Hex(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "lot")
		assert.Nil(err)
		response := ress.(*models.LotModel)
		assert.Equal(lot1.ID, response.ID)
		assert.Equal(lot1.Name, response.Name)
	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not get the lot outside the business", func(t *testing.T) {
			assert := assert.New(t)
			req := helpers.DoRequest("GET", path+"/"+lot4.ID.Hex(), nil, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)
		})
	})

}
