package projects

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/api/projects"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
)

func TestGetProjects(t *testing.T) {
	log.Println("-- test get projects --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/projects"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	project1 := helpers.ProjectFixture(app, token1, 0)
	helpers.ProjectFixture(app, token1, 1)
	project3 := helpers.ProjectFixture(app, token2, 1)

	t.Run("It should get the list of projects", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "projects")
		assert.Nil(err)
		response := ress.(*projects.ResponseList)
		assert.Len(response.Data, 2)
		assert.Equal(response.Total, 2)
	})

	t.Run("It should get single project by ID", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path+"/"+project1.ID.Hex(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "project")
		assert.Nil(err)
		response := ress.(*models.ProjectModel)
		assert.Equal(project1.ID, response.ID)
		assert.Equal(project1.Name, response.Name)
	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not get the project outside the business", func(t *testing.T) {
			assert := assert.New(t)
			req := helpers.DoRequest("GET", path+"/"+project3.ID.Hex(), nil, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)
		})
	})

}
