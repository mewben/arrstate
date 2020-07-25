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

func TestMe(t *testing.T) {
	log.Println("[TEST ME]")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/me"

	// Setup
	helpers.CleanupFixture(db)
	token := helpers.SignupFixture(app, 0)

	t.Run("It should get me details", func(t *testing.T) {
		assert := assert.New(t)
		req := helpers.DoRequest("GET", path, nil, token)
		signupFakeData := helpers.FakeSignup[0]

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "me")
		assert.Nil(err)
		response := ress.(*models.MeModel)
		user := response.CurrentUser.User
		person := response.CurrentUser.Person
		business := response.CurrentBusiness

		assert.False(business.ID.IsZero())
		assert.Equal(signupFakeData.Domain, business.Domain)
		assert.Equal(signupFakeData.Business, business.Name)
		assert.Len(response.UserBusinesses, 1)
		assert.NotNil(person)
		assert.NotNil(user)
		assert.Empty(user.Password)
		assert.Equal(signupFakeData.Email, user.Email)
		assert.Equal(enums.AccountStatusActive, user.AccountStatus)
		assert.Equal(business.ID, person.BusinessID)
		assert.Equal(signupFakeData.GivenName, person.GivenName)
		assert.Equal(signupFakeData.FamilyName, person.FamilyName)
		assert.Contains(person.Role, "owner")
	})
}
