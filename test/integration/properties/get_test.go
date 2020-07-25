package properties

import (
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/api/properties"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestGetProperties(t *testing.T) {
	log.Println("-- test get properties --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/properties"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	project := helpers.ProjectFixture(app, token1, 0)
	project2 := helpers.ProjectFixture(app, token2, 1)
	property1 := helpers.PropertyFixture(app, token1, &project.ID, 0)
	helpers.PropertyFixture(app, token1, &project.ID, 1)
	helpers.PropertyFixture(app, token1, nil, 1)
	property4 := helpers.PropertyFixture(app, token2, &project2.ID, 1)

	t.Run("It should get the list of properties", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "properties")
		assert.Nil(err)
		response := ress.(*properties.ResponseList)
		assert.Len(response.Data, 3)
		assert.Equal(response.Total, 3)
	})

	t.Run("It should get the list of properties with projectID", func(t *testing.T) {
		assert := assert.New(t)
		query := url.Values{}
		query.Add("projectID", project.ID.Hex())
		req := helpers.DoRequest("GET", path+"?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "properties")
		assert.Nil(err)
		response := ress.(*properties.ResponseList)
		assert.Len(response.Data, 2)
		assert.Equal(response.Total, 2)
	})

	t.Run("It should get the list of properties with projectID = null", func(t *testing.T) {
		assert := assert.New(t)
		query := url.Values{}
		query.Add("projectID", "null")
		req := helpers.DoRequest("GET", path+"?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "properties")
		assert.Nil(err)
		response := ress.(*properties.ResponseList)
		assert.Len(response.Data, 1)
		assert.Equal(response.Total, 1)
	})

	t.Run("It should get the property by id", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path+"/"+property1.ID.Hex(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "property")
		assert.Nil(err)
		response := ress.(*models.PropertyModel)
		assert.Equal(property1.ID, response.ID)
		assert.Equal(property1.Name, response.Name)
	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not get the property outside the business", func(t *testing.T) {
			assert := assert.New(t)
			req := helpers.DoRequest("GET", path+"/"+property4.ID.Hex(), nil, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFoundProperty), response.Message, response)
		})
	})

}
