package projects

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

func TestRemoveProject(t *testing.T) {
	log.Println("-- test remove project --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/projects"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	project := helpers.ProjectFixture(app, token1, 0)
	project2 := helpers.ProjectFixture(app, token2, 1)

	t.Run("It should remove project", func(t *testing.T) {
		assert := assert.New(t)
		projectID := project.ID.Hex()
		req := helpers.DoRequest("DELETE", path+"/"+projectID, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseDelete(res)
		assert.Nil(err)
		assert.Equal(projectID, response["project"])

	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not remove project from other business", func(t *testing.T) {
			projectID := project2.ID.Hex()
			assert := assert.New(t)
			req := helpers.DoRequest("DELETE", path+"/"+projectID, nil, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFoundProject), response.Message, response)
		})
	})
}
