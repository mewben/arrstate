package auth

import (
	"context"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/auth"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
)

func TestIntegrationSignin(t *testing.T) {
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/auth/signin"

	// cleanup
	db.Collection(enums.CollBusinesses).DeleteMany(context.Background(), bson.D{{}})
	db.Collection(enums.CollUsers).DeleteMany(context.Background(), bson.D{{}})
	db.Collection(enums.CollPeople).DeleteMany(context.Background(), bson.D{{}})

	// Setup, insert valid business and user first
	payload := &auth.SignupPayload{
		GivenName:  "testgn",
		FamilyName: "testfn",
		Business:   "Test Business",
		Domain:     "test-domain",
		Email:      "test@email.com",
		Password:   "password",
	}
	req := helpers.DoRequest("POST", "/auth/signup", payload)
	app.Test(req, -1)

	t.Run("It should return the JWT and authSuccess data", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"email":    payload.Email,
			"password": payload.Password,
		}
		req := helpers.DoRequest("POST", path, data)
		res, err := app.Test(req, -1)

		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseAuth(res)
		assert.Nil(err)
		user := response.CurrentUser.User
		person := response.CurrentUser.Person
		business := response.CurrentBusiness
		checkJWT(response.Token, user, assert)
		assert.NotEmpty(business.ID)
		assert.Equal(business.Domain, payload.Domain)
		assert.Equal(business.Name, payload.Business)
		assert.Equal(len(response.UserBusinesses), 1)
		assert.NotNil(person)
		assert.NotNil(user)
		assert.Empty(user.Password)
		assert.Equal(user.Email, payload.Email)
		assert.Equal(user.AccountStatus, enums.AccountStatusPending)
		assert.Equal(person.BusinessID, business.ID)
		assert.Equal(person.GivenName, payload.GivenName)
		assert.Equal(person.FamilyName, payload.FamilyName)
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
			req := helpers.DoRequest("POST", path, cas.Payload)
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
