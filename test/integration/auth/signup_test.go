package auth

import (
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

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
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/auth/signup"

	// setup
	helpers.CleanupFixture(db)

	t.Run("It should return the JWT and authSuccess data", func(t *testing.T) {
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
		assert.Equal(201, res.StatusCode, res)
		response, err := helpers.GetResponseAuth(res)
		assert.Nil(err)
		user := response.CurrentUser.User
		person := response.CurrentUser.Person
		business := response.CurrentBusiness
		checkJWT(response.Token, user, business.ID, assert)
		assert.NotEmpty(business.ID)
		assert.Equal(business.Domain, fakeDomain)
		assert.Equal(business.Name, fakeBusiness)
		assert.Equal(len(response.UserBusinesses), 1)
		assert.NotNil(person)
		assert.NotNil(user)
		assert.Empty(user.Password)
		assert.Equal(user.Email, fakeEmail)
		assert.Equal(user.AccountStatus, enums.AccountStatusPending)
		assert.Equal(person.BusinessID, business.ID)
		assert.Equal(person.GivenName, fakeGivenname)
		assert.Equal(person.FamilyName, fakeFamilyName)
		assert.Equal(person.Role, "owner")
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
		response, err := helpers.GetResponseAuth(res)
		assert.Nil(err)
		assert.Equal(response.CurrentBusiness.Domain, "test-domain-2")
		assert.Equal(response.CurrentUser.User.Email, "test2@email.com")
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

func checkJWT(token string, user *models.UserModel, businessID string, assert *assert.Assertions) {
	assert.NotEmpty(token)
	tokenSigningKey := viper.GetString("TOKEN_SIGNING_KEY")
	assert.NotEmpty(tokenSigningKey)
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenSigningKey), nil
	})
	assert.Nil(err, t)
	claims := t.Claims.(jwt.MapClaims)
	exp := time.Now().Add(time.Hour * viper.GetDuration("TOKEN_EXPIRY")).Unix()
	claimsExpiry := claims["exp"].(float64)
	diff := float64(exp) - claimsExpiry
	assert.Equal(user.ID, claims["sub"])
	assert.LessOrEqual(diff, float64(1))
	assert.Equal(businessID, claims["businessID"])
}
