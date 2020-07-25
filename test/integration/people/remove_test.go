package people

import (
	"log"
	"os"
	"testing"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestRemovePerson(t *testing.T) {
	log.Println("-- test remove person --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/people"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	person1 := helpers.PersonFixture(app, token1, 0)
	person2 := helpers.PersonFixture(app, token2, 1)

	t.Run("It should remove person", func(t *testing.T) {
		assert := assert.New(t)
		personID := person1.ID.Hex()
		req := helpers.DoRequest("DELETE", path+"/"+personID, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseDelete(res)
		assert.Nil(err)
		assert.Equal(personID, response["person"])
	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not remove from other business", func(t *testing.T) {
			assert := assert.New(t)
			personID := person2.ID.Hex()
			req := helpers.DoRequest("DELETE", path+"/"+personID, nil, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFoundPerson), response.Message, response)

		})

		t.Run("It should not remove self", func(t *testing.T) {
			// TODO
		})

		t.Run("It should not remove if role not permitted", func(t *testing.T) {
			// TODO
		})
	})
}
