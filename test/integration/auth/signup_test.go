package auth

import (
	"context"
	"log"
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
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
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
		fakeGivenName := "Given Name"
		fakeFamilyName := "Family Name"
		fakePassword := "passworD"
		data := fiber.Map{
			"email":      fakeEmail,
			"password":   fakePassword,
			"business":   fakeBusiness,
			"domain":     fakeDomain,
			"givenName":  fakeGivenName,
			"familyName": fakeFamilyName,
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
		assert.Equal(fakeGivenName, person.GivenName)
		assert.Equal(fakeFamilyName, person.FamilyName)
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
		fakeGivenname := "Given Name"
		fakeFamilyName := "Family Name"
		fakePassword := "passworD"
		data := fiber.Map{
			"email":      fakeEmail,
			"password":   fakePassword,
			"business":   fakeBusiness,
			"domain":     fakeDomain,
			"givenName":  fakeGivenname,
			"familyName": fakeFamilyName,
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
				GivenName:  "",
				FamilyName: "",
				Business:   "",
				Domain:     "",
				Email:      "",
				Password:   "",
			},
			{
				GivenName:  "",
				FamilyName: "testfn",
				Business:   "testb",
				Domain:     "testd",
				Email:      "teste",
				Password:   "test",
			},
			{
				GivenName:  "testgn",
				FamilyName: "testfn",
				Business:   "",
				Domain:     "testd",
				Email:      "teste",
				Password:   "test",
			},
			{
				GivenName:  "testgn",
				FamilyName: "testfn",
				Business:   "testb",
				Domain:     "",
				Email:      "teste",
				Password:   "test",
			},
			{
				GivenName:  "testgn",
				FamilyName: "testfn",
				Business:   "testb",
				Domain:     "testd",
				Email:      "",
				Password:   "test",
			},
			{
				GivenName:  "testgn",
				FamilyName: "testfn",
				Business:   "testb",
				Domain:     "testd",
				Email:      "teste",
				Password:   "",
			},
			{
				// invalid email
				GivenName:  "testgn",
				FamilyName: "testfn",
				Business:   "testb",
				Domain:     "testd",
				Email:      "teste",
				Password:   "test",
			},
			{
				// password must be at least 6 chars
				GivenName:  "testgn",
				FamilyName: "testfn",
				Business:   "testb",
				Domain:     "testd",
				Email:      "sample@email.com",
				Password:   "test",
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
		fakeGivenname := "Given Name"
		fakeFamilyName := "Family Name"
		fakePassword := "passworD"
		data := fiber.Map{
			"email":      fakeEmail,
			"password":   fakePassword,
			"business":   fakeBusiness,
			"domain":     fakeDomain,
			"givenName":  fakeGivenname,
			"familyName": fakeFamilyName,
		}
		req := helpers.DoRequest("POST", path, data, "")

		// Execute -
		res, err := app.Test(req, -1)
		// Assert -
		assert.Nil(err)
		assert.Equal(400, res.StatusCode, res)
		response, err := helpers.GetResponseError(res)
		assert.Nil(err)
		assert.Equal(response.Message, services.T(errors.ErrDomainDuplicate), response)

	})

	t.Run("It should throw duplicate email", func(t *testing.T) {
		// Setup -
		assert := assert.New(t)
		fakeEmail := "test@email.com"
		fakeBusiness := "Test Business"
		fakeDomain := "test-domain-3"
		fakeGivenname := "Given Name"
		fakeFamilyName := "Family Name"
		fakePassword := "passworD"
		data := fiber.Map{
			"email":      fakeEmail,
			"password":   fakePassword,
			"business":   fakeBusiness,
			"domain":     fakeDomain,
			"givenName":  fakeGivenname,
			"familyName": fakeFamilyName,
		}
		req := helpers.DoRequest("POST", path, data, "")

		// Execute -
		res, err := app.Test(req, -1)
		// Assert -
		assert.Nil(err)
		assert.Equal(400, res.StatusCode, res)
		response, err := helpers.GetResponseError(res)
		assert.Nil(err)
		assert.Equal(response.Message, services.T(errors.ErrUserDuplicate), response)
	})

	t.Run("It should cleanup business, user, people on signup error", func(t *testing.T) {
		// TODO: not urgent
	})

}
