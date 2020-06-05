package projects

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/test/helpers"
)

func TestGetProject(t *testing.T) {
	log.Println("-- test get projects --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/projects"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app, 1)
	_, authResponse2 := helpers.SignupFixture(app, 2)
	helpers.ProjectFixture(app, authResponse.Token, 1)
	helpers.ProjectFixture(app, authResponse.Token, 2)
	helpers.ProjectFixture(app, authResponse2.Token, 2)

	t.Run("It should get the list of projects", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path, nil, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseProjects(res)
		assert.Nil(err)
		assert.Len(response.Data, 2)
		assert.Equal(response.Total, 2)
	})

}
