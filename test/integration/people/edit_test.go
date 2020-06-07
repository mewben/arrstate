package people

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
)

func TestEditPerson(t *testing.T) {
	log.Println("-- test edit person --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/people"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app, 1)
	_, authResponse2 := helpers.SignupFixture(app, 2)
	helpers.PersonFixture(app, authResponse.Token, 1) // person1
	person2 := helpers.PersonFixture(app, authResponse2.Token, 1)

	t.Run("It should edit person", func(t *testing.T) {
		assert := assert.New(t)
		updEmail := "editemail@email.com"
		updRole := enums.RoleClient
		updGivenName := "edit1"
		updFamilyName := "edit2"
		updAddress := models.NewAddressModel()
		updAddress.Country = "US"
		updAddress.State = "Ohio"
		updNotes := "Edit notes"
		updCommissionPerc := 14.5
		updCustom := fiber.Map{
			"tin":     "editin",
			"contact": "editcontact",
			"color":   "red",
		}
		data := fiber.Map{
			"_id":            authResponse.CurrentUser.Person.ID,
			"email":          updEmail,
			"role":           updRole,
			"givenName":      updGivenName,
			"familyName":     updFamilyName,
			"address":        updAddress,
			"notes":          updNotes,
			"commissionPerc": updCommissionPerc,
			"customFields":   updCustom,
		}
		req := helpers.DoRequest("PUT", path, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		response, err := helpers.GetResponsePerson(res)
		assert.Nil(err)
		assert.Equal(authResponse.CurrentBusiness.ID, response.BusinessID)
		assert.Equal(authResponse.CurrentUser.User.ID, response.UpdatedBy)
		assert.Nil(response.UserID)
		assert.False(response.ID.IsZero())
		assert.Equal(updRole, response.Role)
		assert.Equal(updGivenName, response.GivenName)
		assert.Equal(updFamilyName, response.FamilyName)
		assert.Equal(updAddress.Country, response.Address.Country)
		assert.Equal(updAddress.State, response.Address.State)
		assert.Equal(updNotes, response.Notes)
		assert.EqualValues(updCommissionPerc, response.CommissionPerc)
		assert.Equal(updCustom, response.CustomFields)

	})

	t.Run("It should validate input", func(t *testing.T) {
		assert := assert.New(t)
		cases := []struct {
			err     string
			payload fiber.Map
		}{
			{
				errors.ErrInputInvalid,
				fiber.Map{},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":  authResponse.CurrentUser.Person.ID,
					"role": "",
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":       authResponse.CurrentUser.Person.ID,
					"role":      enums.RoleAgent,
					"givenName": "",
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":       authResponse.CurrentUser.Person.ID,
					"role":      enums.RoleAgent,
					"givenName": "givenName",
					"email":     "",
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":       authResponse.CurrentUser.Person.ID,
					"role":      enums.RoleAgent,
					"givenName": "givenName",
					"email":     "notavalidemail",
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":            authResponse.CurrentUser.Person.ID,
					"role":           enums.RoleAgent,
					"givenName":      "givenName",
					"email":          "sample@email.com",
					"commissionPerc": -34,
				},
			},
			{
				errors.ErrUpdate,
				fiber.Map{
					// duplicate person for this business
					"_id":            authResponse.CurrentUser.Person.ID,
					"role":           enums.RoleAgent,
					"givenName":      "givenNameerr",
					"email":          "email1@email.com",
					"commissionPerc": 34,
				},
			},
		}

		for _, item := range cases {
			req := helpers.DoRequest("PUT", path, item.payload, authResponse.Token)
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, item.payload)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(item.err), response.Message, response)
		}

	})

	t.Run("It should not be able to edit role as owner", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"_id":       person2.ID,
			"role":      enums.RoleOwner,
			"email":     "eddfe@email.com",
			"givenName": "given",
		}
		req := helpers.DoRequest("PUT", path, data, authResponse2.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(400, res.StatusCode, res)
		response, err := helpers.GetResponseError(res)
		assert.Nil(err)
		assert.Equal(services.T(errors.ErrOwner), response.Message, response)
	})

	t.Run("Permissions", func(t *testing.T) {

		t.Run("It should not edit from other business", func(t *testing.T) {
			assert := assert.New(t)
			data := fiber.Map{
				"_id":       person2.ID,
				"role":      enums.RoleAgent,
				"email":     "eddfe@email.com",
				"givenName": "given",
			}
			req := helpers.DoRequest("PUT", path, data, authResponse.Token)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFound), response.Message, response)

		})
	})
}
