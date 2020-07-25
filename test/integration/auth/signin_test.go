package auth

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/auth"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
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
		signupFakeData := helpers.FakeSignup[0]
		data := fiber.Map{
			"email":    signupFakeData.Email,
			"password": signupFakeData.Password,
		}
		req := helpers.DoRequest("POST", path, data, "")
		req.Header.Add("origin", "http://test-domain.example.com")
		res, err := app.Test(req, -1)

		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		helpers.CheckJWT(response["token"].(string), assert)
	})

	t.Run("It should login using deviceCode and remove it after", func(t *testing.T) {
		assert := assert.New(t)
		// signup
		req := helpers.DoRequest("POST", "/auth/signup", helpers.FakeSignup[1], "")
		res, err := app.Test(req, -1)
		assert.Nil(err)
		response, err := helpers.GetResponseMap(res)
		assert.Nil(err)

		signinPayload := fiber.Map{
			"grant_type": "device_code",
			"deviceCode": response["deviceCode"],
		}
		req = helpers.DoRequest("POST", path, signinPayload, "")
		req.Header.Add("origin", "http://test-domain2.example.com")
		res, err = app.Test(req, -1)
		assert.Nil(err)
		response, err = helpers.GetResponseMap(res)
		assert.Nil(err)
		userID, _ := helpers.CheckJWT(response["token"].(string), assert)

		// get user
		filter := bson.D{
			{
				Key:   "_id",
				Value: userID,
			},
		}
		user := &models.UserModel{}
		// delay to complete go func in deleting the deviceCode
		time.Sleep(time.Second * 1)
		err = db.Collection(enums.CollUsers).FindOne(context.TODO(), filter).Decode(&user)
		assert.Nil(err)
		assert.Empty(user.DeviceCode)

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
