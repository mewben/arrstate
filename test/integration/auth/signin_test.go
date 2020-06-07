package auth

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/auth"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
)

func TestSignin(t *testing.T) {
	log.Println("[TEST SIGNIN]")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/auth/signin"

	// Setup
	helpers.CleanupFixture(db)
	signupPayload, _ := helpers.SignupFixture(app, 1)

	t.Run("It should return the JWT and authSuccess data", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"email":    signupPayload.Email,
			"password": signupPayload.Password,
		}
		req := helpers.DoRequest("POST", path, data, "")
		res, err := app.Test(req, -1)

		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseAuth(res)
		assert.Nil(err)
		user := response.CurrentUser.User
		person := response.CurrentUser.Person
		business := response.CurrentBusiness
		helpers.CheckJWT(response.Token, user, business.ID, assert)
		assert.False(business.ID.IsZero())
		assert.Equal(signupPayload.Domain, business.Domain)
		assert.Equal(signupPayload.Business, business.Name)
		assert.Len(response.UserBusinesses, 1)
		assert.NotNil(person)
		assert.NotNil(user)
		assert.Empty(user.Password)
		assert.Equal(signupPayload.Email, user.Email)
		assert.Equal(enums.AccountStatusPending, user.AccountStatus)
		assert.Equal(business.ID, person.BusinessID)
		assert.Equal(signupPayload.GivenName, person.GivenName)
		assert.Equal(signupPayload.FamilyName, person.FamilyName)
		assert.Equal("owner", person.Role)
	})

	t.Run("It should catch invalid email or password", func(t *testing.T) {
		assert := assert.New(t)
		cases := []struct {
			Err     string
			Payload auth.SigninPayload
		}{
			{
				errors.ErrInputInvalid,
				auth.SigninPayload{
					Email:    "",
					Password: "",
				},
			},
			{
				errors.ErrInputInvalid,
				auth.SigninPayload{
					Email:    "test@email.com",
					Password: "",
				},
			},
			{
				errors.ErrInputInvalid,
				auth.SigninPayload{
					Email:    "test@email.com",
					Password: "short",
				},
			},
			{
				errors.ErrSigninIncorrect,
				auth.SigninPayload{
					Email:    "test@email.com",
					Password: "wrongpassword",
				},
			},
			{
				errors.ErrInputInvalid,
				auth.SigninPayload{
					Email:    "",
					Password: "password",
				},
			},
			{
				errors.ErrSigninIncorrect,
				auth.SigninPayload{
					Email:    "wrong@email.com",
					Password: "password",
				},
			},
		}

		for _, cas := range cases {
			req := helpers.DoRequest("POST", path, cas.Payload, "")
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(cas.Err), response.Message, response)
		}

	})

	t.Run("It should lock for 5 minutes if 5 unsuccessful attempts", func(t *testing.T) {
		// TODO: not urgent
	})

	// startup.DisconnectMongo(db)
}
