package users

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/test/helpers"
)

func TestGetCurrent(t *testing.T) {
	log.Println("GET CURRENT USER")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/users/current"

	// Setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	userID, _ := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should get the details of the current user", func(t *testing.T) {
		assert := assert.New(t)
		signupData := helpers.FakeSignup[0]
		req := helpers.DoRequest("GET", path, nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "user")
		assert.Nil(err)
		response := ress.(*models.UserModel)
		assert.Equal(userID, response.ID)
		assert.Empty(response.Password)
		assert.Empty(response.DeviceCode)
		assert.Equal(enums.AccountStatusActive, response.AccountStatus)
		assert.Equal(signupData.Email, response.Email)
	})

}
