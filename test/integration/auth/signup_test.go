package auth

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
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

func TestSignup(t *testing.T) {
	log.Println("[TEST SIGNUP]")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/auth/signup"

	// setup
	helpers.CleanupFixture(db)

	t.Run("It should return the DeviceCode and domain", func(t *testing.T) {
		// Setup -
		assert := assert.New(t)
		fakeEmail := "test@email.com"
		fakeBusiness := "Test Business"
		fakeDomain := "test-domain"
		fakeFirstName := "Given Name"
		fakeLastName := "Family Name"
		fakePassword := "passworD"
		data := fiber.Map{
			"email":    fakeEmail,
			"password": fakePassword,
			"business": fakeBusiness,
			"domain":   fakeDomain,
			"name": fiber.Map{
				"first": fakeFirstName,
				"last":  fakeLastName,
			},
		}
		req := helpers.DoRequest("POST", path, data, "")

		// Execute -
		res, err := app.Test(req, -1)
		// Assert -
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		response, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		assert.NotEmpty(response["deviceCode"])
		// userID, businessID := helpers.CheckJWT(response["token"].(string), assert)

		// get business
		filter := bson.D{
			{
				Key:   "domain",
				Value: response["domain"],
			},
		}
		business := &models.BusinessModel{}
		err = db.Collection(enums.CollBusinesses).FindOne(context.TODO(), filter).Decode(&business)
		assert.Nil(err)
		assert.Equal(fakeBusiness, business.Name)
		assert.Equal(fakeDomain, business.Domain)

		// get person
		filter = bson.D{
			{
				Key:   "businessID",
				Value: business.ID,
			},
		}
		person := &models.PersonModel{}
		err = db.Collection(enums.CollPeople).FindOne(context.TODO(), filter).Decode(&person)
		assert.Nil(err)
		assert.Equal(fakeFirstName, person.Name.First)
		assert.Equal(fakeLastName, person.Name.Last)
		assert.Equal(enums.DefaultDateFormat, person.Locale.DateFormat)
		assert.Equal(enums.DefaultTimeFormat, person.Locale.TimeFormat)
		assert.Equal(enums.DefaultTimestampFormat, person.Locale.TimestampFormat)
		assert.Equal("sunday", person.Locale.WeekStartDay)

		// get user
		filter = bson.D{
			{
				Key:   "_id",
				Value: person.UserID,
			},
		}
		user := &models.UserModel{}
		err = db.Collection(enums.CollUsers).FindOne(context.TODO(), filter).Decode(&user)
		assert.Nil(err)
		assert.Equal(fakeEmail, user.Email)
		assert.NotEqual(fakePassword, user.Password)
		assert.NotEmpty(user.DeviceCode)

	})

	t.Run("It should auto clean inputs", func(t *testing.T) {
		// Setup -
		assert := assert.New(t)
		fakeEmail := "Test2@email.com"
		fakeBusiness := "Test Business2"
		fakeDomain := "Test Domain 2"
		fakeFirstName := "Given Name"
		fakeLastName := "Family Name"
		fakePassword := "passworD"
		data := fiber.Map{
			"email":    fakeEmail,
			"password": fakePassword,
			"business": fakeBusiness,
			"domain":   fakeDomain,
			"name": fiber.Map{
				"first": fakeFirstName,
				"last":  fakeLastName,
			},
		}
		req := helpers.DoRequest("POST", path, data, "")

		// Execute -
		res, err := app.Test(req, -1)
		// Assert -
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		response, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		// userID, businessID := helpers.CheckJWT(response["token"].(string), assert)

		// get business
		filter := bson.D{
			{
				Key:   "domain",
				Value: response["domain"],
			},
		}
		business := &models.BusinessModel{}
		err = db.Collection(enums.CollBusinesses).FindOne(context.TODO(), filter).Decode(&business)
		assert.Nil(err)
		assert.Equal("test-domain-2", business.Domain)

		// get person
		filter = bson.D{
			{
				Key:   "businessID",
				Value: business.ID,
			},
		}
		person := &models.PersonModel{}
		err = db.Collection(enums.CollPeople).FindOne(context.TODO(), filter).Decode(&person)
		assert.Nil(err)
		assert.NotNil(person)

		// get user
		filter = bson.D{
			{
				Key:   "_id",
				Value: person.UserID,
			},
		}
		user := &models.UserModel{}
		err = db.Collection(enums.CollUsers).FindOne(context.TODO(), filter).Decode(&user)
		assert.Nil(err)
		assert.Equal("test2@email.com", user.Email)

	})

	t.Run("It should validate inputs", func(t *testing.T) {
		assert := assert.New(t)
		payloads := []auth.SignupPayload{
			{},
			{
				Name: models.PersonName{
					First: "",
					Last:  "",
				},
				Business: "",
				Domain:   "",
				Email:    "",
				Password: "",
			},
			{
				Name: models.PersonName{
					First: "",
					Last:  "testlast",
				},
				Business: "testb",
				Domain:   "testd",
				Email:    "teste",
				Password: "test",
			},
			{
				Name: models.PersonName{
					First: "testfn",
					Last:  "testln",
				},
				Business: "",
				Domain:   "testd",
				Email:    "teste",
				Password: "test",
			},
			{
				Name: models.PersonName{
					First: "testfn",
					Last:  "testln",
				},
				Business: "testb",
				Domain:   "",
				Email:    "teste",
				Password: "test",
			},
			{
				Name: models.PersonName{
					First: "testfn",
					Last:  "testln",
				},
				Business: "testb",
				Domain:   "testd",
				Email:    "",
				Password: "test",
			},
			{
				Name: models.PersonName{
					First: "testfn",
					Last:  "testln",
				},
				Business: "testb",
				Domain:   "testd",
				Email:    "teste",
				Password: "",
			},
			{
				// invalid email
				Name: models.PersonName{
					First: "testfn",
					Last:  "testln",
				},
				Business: "testb",
				Domain:   "testd",
				Email:    "teste",
				Password: "test",
			},
			{
				// password must be at least 6 chars
				Name: models.PersonName{
					First: "testfn",
					Last:  "testln",
				},
				Business: "testb",
				Domain:   "testd",
				Email:    "sample@email.com",
				Password: "test",
			},
		}

		for _, payload := range payloads {
			req := helpers.DoRequest("POST", path, payload, "")
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(response.Message, services.T(errors.ErrInputInvalid), response)
		}

	})

	t.Run("It should throw duplicate domain", func(t *testing.T) {
		// Setup -
		assert := assert.New(t)
		fakeEmail := "test@email.com"
		fakeBusiness := "Test Business"
		fakeDomain := "test-domain"
		fakeFirstName := "Given Name"
		fakeLastName := "Family Name"
		fakePassword := "passworD"
		data := fiber.Map{
			"email":    fakeEmail,
			"password": fakePassword,
			"business": fakeBusiness,
			"domain":   fakeDomain,
			"name": fiber.Map{
				"first": fakeFirstName,
				"last":  fakeLastName,
			},
		}
		req := helpers.DoRequest("POST", path, data, "")

		// Execute -
		res, err := app.Test(req, -1)
		// Assert -
		assert.Nil(err)
		assert.Equal(400, res.StatusCode, res)
		response, err := helpers.GetResponseError(res)
		assert.Nil(err)
		assert.Equal(response.Message, services.T(errors.ErrDuplicateDomain), response)

	})

	t.Run("It should throw duplicate email", func(t *testing.T) {
		// Setup -
		assert := assert.New(t)
		fakeEmail := "test@email.com"
		fakeBusiness := "Test Business"
		fakeDomain := "test-domain-3"
		fakeFirstName := "Given Name"
		fakeLastName := "Family Name"
		fakePassword := "passworD"
		data := fiber.Map{
			"email":    fakeEmail,
			"password": fakePassword,
			"business": fakeBusiness,
			"domain":   fakeDomain,
			"name": fiber.Map{
				"first": fakeFirstName,
				"last":  fakeLastName,
			},
		}
		req := helpers.DoRequest("POST", path, data, "")

		// Execute -
		res, err := app.Test(req, -1)
		// Assert -
		assert.Nil(err)
		assert.Equal(400, res.StatusCode, res)
		response, err := helpers.GetResponseError(res)
		assert.Nil(err)
		assert.Equal(response.Message, services.T(errors.ErrDuplicateUser), response)
	})

	t.Run("It should cleanup business, user, people on signup error", func(t *testing.T) {
		// TODO: not urgent
	})

}
