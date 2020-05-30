package auth

import (
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/test/helpers"
)

func TestSignup(t *testing.T) {
	os.Setenv("ENV", "TESTING")
	internal.InitEnvironment()

	path := "/auth/signup"

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
			"email":        fakeEmail,
			"password":     fakePassword,
			"businessName": fakeBusiness,
			"domain":       fakeDomain,
			"givenName":    fakeGivenname,
			"familyName":   fakeFamilyName,
		}
		h := &Handler{}
		app := fiber.New()
		app.Post(path, h.Signup)
		req := helpers.DoRequest("POST", path, data)

		// Execute -
		res, err := app.Test(req, -1)
		// Assert -
		assert.Nil(err)
		assert.Equal(200, res.StatusCode)
		response, err := helpers.GetResponseAuth(res)
		assert.Nil(err)
		user := response.CurrentUser.User
		person := response.CurrentUser.Person
		business := response.CurrentBusiness
		checkJWT(response.Token, user, assert)
		assert.NotEmpty(business.ID)
		assert.Equal(business.Domain, fakeDomain)
		assert.Equal(business.Name, fakeBusiness)
		assert.Equal(len(response.UserBusinesses), 1)
		assert.NotNil(person)
		assert.NotNil(user)
		assert.Empty(user.Password)
		assert.Equal(user.Email, fakeEmail)
		assert.Equal(user.AccountStatus, "pending")
		assert.Equal(person.BusinessID, business.ID)
		assert.Equal(person.GivenName, fakeGivenname)
		assert.Equal(person.FamilyName, fakeFamilyName)
		assert.Equal(person.Role, "owner")
	})

}

func checkJWT(token string, user *models.UserModel, assert *assert.Assertions) {
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
	assert.Equal(diff, float64(0))
}
