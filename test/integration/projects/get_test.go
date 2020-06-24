package projects

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/api/projects"
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
	helpers.ProjectFixture(app, token1, 0)
	helpers.ProjectFixture(app, token1, 1)
	helpers.ProjectFixture(app, token2, 1)

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

}
