package lots

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
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
	_, authResponse := helpers.SignupFixture(app, 1)
	_, authResponse2 := helpers.SignupFixture(app, 2)
	project := helpers.ProjectFixture(app, authResponse.Token, 1)
	project2 := helpers.ProjectFixture(app, authResponse2.Token, 2)
	helpers.LotFixture(app, authResponse.Token, project.ID, 1)
	helpers.LotFixture(app, authResponse.Token, project.ID, 2)
	helpers.LotFixture(app, authResponse2.Token, project2.ID, 2)

	t.Run("It should get the list of lots", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path+"/"+project.ID.Hex(), nil, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseLots(res)
		assert.Nil(err)
		assert.Len(response.Data, 2)
		assert.Equal(response.Total, 2)
	})

}
