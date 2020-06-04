package auth

import (
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
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/auth/signin"

	// Setup
	helpers.CleanupFixture(db)
	signupPayload, _ := helpers.SignupFixture(app)

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
		assert.NotEmpty(business.ID)
		assert.Equal(business.Domain, signupPayload.Domain)
		assert.Equal(business.Name, signupPayload.Business)
		assert.Equal(len(response.UserBusinesses), 1)
		assert.NotNil(person)
		assert.NotNil(user)
		assert.Empty(user.Password)
		assert.Equal(user.Email, signupPayload.Email)
		assert.Equal(user.AccountStatus, enums.AccountStatusPending)
		assert.Equal(person.BusinessID, business.ID)
		assert.Equal(person.GivenName, signupPayload.GivenName)
		assert.Equal(person.FamilyName, signupPayload.FamilyName)
		assert.Equal(person.Role, "owner")
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
			assert.Equal(response.Message, services.T(cas.Err), response)
		}

	})

	t.Run("It should lock for 5 minutes if 5 unsuccessful attempts", func(t *testing.T) {
		// TODO: not urgent
	})
}
