package people

import (
	"log"
	"os"
	"testing"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
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
	_, authResponse := helpers.SignupFixture(app, 1)
	_, authResponse2 := helpers.SignupFixture(app, 2)
	person1 := helpers.PersonFixture(app, authResponse.Token, 1)
	person2 := helpers.PersonFixture(app, authResponse2.Token, 2)

	t.Run("It should remove person", func(t *testing.T) {
		assert := assert.New(t)
		personID := person1.ID.Hex()
		req := helpers.DoRequest("DELETE", path+"/"+personID, nil, authResponse.Token)

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
			req := helpers.DoRequest("DELETE", path+"/"+personID, nil, authResponse.Token)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)

		})

		t.Run("It should not remove self", func(t *testing.T) {
			// TODO
		})

		t.Run("It should not remove if role not permitted", func(t *testing.T) {
			// TODO
		})
	})
}
