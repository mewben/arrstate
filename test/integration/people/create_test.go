package people

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	log.Println("-- test create lot --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/people"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app, 1)

	t.Run("It should create person inside business", func(t *testing.T) {
		assert := assert.New(t)
		fakeEmail := "test3@email.com"
		fakeRole := enums.RoleCoOwner
		fakeGivenName := "given"
		fakeFamilyName := "family"
		fakeAddress := models.NewAddressModel()
		fakeAddress.Country = "PH"
		fakeAddress.State = "Bohol"
		fakeNotes := "Sample Notes"
		fakeCommissionPerc := 43.2
		fakeCustom := fiber.Map{
			"tin":     "tinno",
			"contact": "contact",
		}
		data := fiber.Map{
			"email":          fakeEmail,
			"role":           fakeRole,
			"givenName":      fakeGivenName,
			"familyName":     fakeFamilyName,
			"address":        fakeAddress,
			"notes":          fakeNotes,
			"commissionPerc": fakeCommissionPerc,
			"customFields":   fakeCustom,
		}
		req := helpers.DoRequest("POST", path, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		response, err := helpers.GetResponsePerson(res)
		assert.Nil(err)
		assert.Equal(authResponse.CurrentBusiness.ID, response.BusinessID)
		assert.Equal(authResponse.CurrentUser.User.ID, response.CreatedBy)
		assert.Nil(response.UserID)
		assert.False(response.ID.IsZero())
		assert.Equal(fakeRole, response.Role)
		assert.Equal(fakeGivenName, response.GivenName)
		assert.Equal(fakeFamilyName, response.FamilyName)
		assert.Equal(fakeAddress.Country, response.Address.Country)
		assert.Equal(fakeAddress.State, response.Address.State)
		assert.Equal(fakeNotes, response.Notes)
		assert.EqualValues(fakeCommissionPerc, response.CommissionPerc)
		assert.Equal(fakeCustom, response.CustomFields)

	})

	t.Run("It should validate input", func(t *testing.T) {
		assert := assert.New(t)
		payloads := []fiber.Map{
			{},
			{
				"role": "",
			},
			{
				"role":      enums.RoleAgent,
				"givenName": "",
			},
			{
				"role":      enums.RoleAgent,
				"givenName": "givenName",
				"email":     "",
			},
			{
				"role":      enums.RoleAgent,
				"givenName": "givenName",
				"email":     "notavalidemail",
			},
			{
				"role":           enums.RoleAgent,
				"givenName":      "givenName",
				"email":          "sample@email.com",
				"commissionPerc": -34,
			},
			{
				// duplicate person for this business
				"role":           enums.RoleAgent,
				"givenName":      "givenName",
				"email":          "test3@email.com",
				"commissionPerc": -34,
			},
		}

		for _, payload := range payloads {
			req := helpers.DoRequest("POST", path, payload, authResponse.Token)
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, payload)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrInputInvalid), response.Message, response)

		}

	})

	t.Run("It should not allow create with owner role", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"email":     "random1@email.com",
			"role":      enums.RoleOwner,
			"givenName": "givenname",
		}
		req := helpers.DoRequest("POST", path, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(400, res.StatusCode, res)
		response, err := helpers.GetResponseError(res)
		assert.Nil(err)
		assert.Equal(services.T(errors.ErrOwner), response.Message, response)
	})
}
