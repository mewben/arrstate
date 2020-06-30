package people

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/api/people"
	"github.com/mewben/realty278/test/helpers"
)

func TestGetPeople(t *testing.T) {
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
	helpers.PersonFixture(app, token1, 2)
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
		assert.Len(response.Data, 4)
		assert.Equal(response.Total, 4)
	})

	t.Run("It should get the list of people by role", func(t *testing.T) {
		assert := assert.New(t)
		query := url.Values{}
		query.Add("role", enums.RoleClient)
		req := helpers.DoRequest("GET", path+"?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "people")
		assert.Nil(err)
		response := ress.(*people.ResponseList)
		assert.Len(response.Data, 2)
		assert.Equal(response.Total, 2)

		// multiple role
		query2 := url.Values{}
		query2.Add("role", enums.RoleClient)
		query2.Add("role", enums.RoleAgent)
		req = helpers.DoRequest("GET", path+"?"+query2.Encode(), nil, token1)

		res, err = app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err = helpers.GetResponse(res, "people")
		assert.Nil(err)
		response = ress.(*people.ResponseList)
		assert.Len(response.Data, 3)
		assert.Equal(response.Total, 3)

	})

}
