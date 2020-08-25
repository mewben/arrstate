package people

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestEditPerson(t *testing.T) {
	log.Println("-- test edit person --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/people"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	person1 := helpers.PersonFixture(app, token1, 0) // person1
	person2 := helpers.PersonFixture(app, token2, 1)
	userID, businessID, _ := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should edit person", func(t *testing.T) {
		assert := assert.New(t)
		updEmail := "editemail@email.com"
		updRole := []string{enums.RoleClient}
		updGivenName := "edit1"
		updFamilyName := "edit2"
		updAddress := models.NewAddressModel()
		updAddress.Country = "US"
		updAddress.State = "Ohio"
		updNotes := "Edit notes"
		updCommissionPerc := 1450 // 14.5%
		updCustom := fiber.Map{
			"tin":     "editin",
			"contact": "editcontact",
			"color":   "red",
		}
		data := fiber.Map{
			"_id":            person1.ID,
			"email":          updEmail,
			"role":           updRole,
			"givenName":      updGivenName,
			"familyName":     updFamilyName,
			"address":        updAddress,
			"notes":          updNotes,
			"commissionPerc": updCommissionPerc,
			"customFields":   updCustom,
		}
		req := helpers.DoRequest("PUT", path, data, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "person")
		assert.Nil(err)
		response := ress.(*models.PersonModel)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.UpdatedBy)
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
					"_id":  person1.ID,
					"role": "",
					"i":    1,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":       person1.ID,
					"role":      []string{enums.RoleAgent},
					"givenName": "",
					"i":         2,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":       person1.ID,
					"role":      []string{enums.RoleAgent},
					"givenName": "givenName",
					"email":     "",
					"i":         3,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":       person1.ID,
					"role":      []string{enums.RoleAgent},
					"givenName": "givenName",
					"email":     "notavalidemail",
					"i":         4,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"_id":            person1.ID,
					"role":           []string{enums.RoleAgent},
					"givenName":      "givenName",
					"email":          "sample@email.com",
					"commissionPerc": -34,
					"i":              5,
				},
			},
			{
				errors.ErrUpdate,
				fiber.Map{
					// duplicate person for this business
					"_id":            person1.ID,
					"role":           []string{enums.RoleAgent},
					"givenName":      "givenNameerr",
					"email":          "test@email.com",
					"commissionPerc": 34,
					"i":              6,
				},
			},
		}

		for _, item := range cases {
			req := helpers.DoRequest("PUT", path, item.payload, token1)
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
			"role":      []string{enums.RoleOwner},
			"email":     "eddfe@email.com",
			"givenName": "given",
		}
		req := helpers.DoRequest("PUT", path, data, token2)

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
				"role":      []string{enums.RoleAgent},
				"email":     "eddfe@email.com",
				"givenName": "given",
			}
			req := helpers.DoRequest("PUT", path, data, token1)

			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFoundPerson), response.Message, response)

		})
	})
}
