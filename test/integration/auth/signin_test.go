package auth

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

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
	helpers.SignupFixture(app, 0)

	t.Run("It should return the JWT", func(t *testing.T) {
		assert := assert.New(t)
		signupFakeData := helpers.SignupFakeData[0]
		data := fiber.Map{
			"email":    signupFakeData.Email,
			"password": signupFakeData.Password,
		}
		req := helpers.DoRequest("POST", path, data, "")
		res, err := app.Test(req, -1)

		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		helpers.CheckJWT(response["token"].(string), assert)
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
