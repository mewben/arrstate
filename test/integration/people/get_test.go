package people

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/api/people"
	"github.com/mewben/realty278/test/helpers"
)

func TestGetPeople(t *testing.T) {
	log.Println("-- test remove person --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/people"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	helpers.PersonFixture(app, token1, 0)
	helpers.PersonFixture(app, token1, 1)
	helpers.PersonFixture(app, token2, 1)

	t.Run("It should get the list of people inside the business", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "people")
		assert.Nil(err)
		response := ress.(*people.ResponseList)
		assert.Len(response.Data, 3)
		assert.Equal(response.Total, 3)
	})

}
