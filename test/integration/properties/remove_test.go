package properties

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestRemoveProperty(t *testing.T) {
	log.Println("-- test remove property --")

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
	property2 := helpers.PropertyFixture(app, token2, &project2.ID, 1)

	t.Run("It should remove property", func(t *testing.T) {
		propertyID := property1.ID.Hex()
		assert := assert.New(t)
		req := helpers.DoRequest("DELETE", path+"/"+propertyID, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseDelete(res)
		assert.Nil(err)
		assert.Equal(propertyID, response["property"])

	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not remove property from other business", func(t *testing.T) {
			propertyID := property2.ID.Hex()
			assert := assert.New(t)
			req := helpers.DoRequest("DELETE", path+"/"+propertyID, nil, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFoundProperty), response.Message, response)
		})
	})
}
