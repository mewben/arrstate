package lots

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/api/lots"
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
	helpers.LotFixture(app, token1, project.ID, 0)
	helpers.LotFixture(app, token1, project.ID, 1)
	helpers.LotFixture(app, token2, project2.ID, 1)

	t.Run("It should get the list of lots", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path+"/"+project.ID.Hex(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "lots")
		assert.Nil(err)
		response := ress.(*lots.ResponseList)
		assert.Len(response.Data, 2)
		assert.Equal(response.Total, 2)
	})

}
