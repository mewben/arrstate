package properties

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestAcquireProperty(t *testing.T) {
	log.Println("-- test acquire property --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/properties/acquire"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	project1 := helpers.ProjectFixture(app, token1, 0)
	project2 := helpers.ProjectFixture(app, token2, 0)
	property1 := helpers.PropertyFixture(app, token1, &project1.ID, 0)
	property2 := helpers.PropertyFixture(app, token1, &project1.ID, 0)
	property3 := helpers.PropertyFixture(app, token1, &project1.ID, 0)
	property4 := helpers.PropertyFixture(app, token2, &project2.ID, 0)
	person1 := helpers.PersonFixture(app, token1, 0)
	person2 := helpers.PersonFixture(app, token1, 1)
	// person3 := helpers.PersonFixture(app, token2, 1)
	userID, businessID := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should let a client acquire a property in cash", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"propertyID":    property1.ID,
			"clientID":      person1.ID,
			"paymentScheme": enums.PaymentSchemeCash,
			"agentID":       person2.ID,
		}
		req := helpers.DoRequest("POST", path, data, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "property")
		assert.Nil(err)
		response := ress.(*models.PropertyModel)
		assert.Equal(enums.StatusAcquired, response.Status)
		// assert acquisition
		acquisition := response.Acquisition
		assert.Equal(person1.ID, *acquisition.ClientID)
		assert.Equal(person2.ID, *acquisition.AgentID)
		assert.Equal(enums.PaymentSchemeCash, acquisition.PaymentScheme)
		assert.Empty(acquisition.PaymentPeriod)
		assert.Empty(acquisition.Terms)
		assert.NotNil(acquisition.AcquiredAt)
		assert.Equal(acquisition.AcquiredAt, acquisition.CompletedAt)

		// assert created invoices
		invoices := make([]*models.InvoiceModel, 0)
		cursor, err := db.Collection(enums.CollInvoices).Find(context.TODO(), bson.M{"propertyID": property1.ID})
		assert.Nil(err)
		err = cursor.All(context.TODO(), &invoices)
		assert.Nil(err)
		assert.Len(invoices, 1)
		invoice := invoices[0]
		assert.Equal(businessID, *invoice.From.ID)
		assert.Equal(userID, invoice.CreatedBy)
		assert.EqualValues(10000050, invoice.Total)
		assert.Equal(fmt.Sprintf("%s - 0", property1.Name), invoice.Name)
		assert.EqualValues(1, invoice.No)

		// TODO: assert created agent commissions

	})

	t.Run("It should let a client acquire a property in installment", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"propertyID":    property2.ID,
			"clientID":      person1.ID,
			"paymentScheme": enums.PaymentSchemeInstallment,
			"paymentPeriod": enums.PaymentPeriodMonthly,
			"terms":         12, // 12 months
			"downPayment":   10000,
		}
		req := helpers.DoRequest("POST", path, data, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "property")
		assert.Nil(err)
		response := ress.(*models.PropertyModel)
		assert.Equal(enums.StatusOnGoing, response.Status)
		// assert acquisition
		acquisition := response.Acquisition
		assert.Equal(person1.ID, *acquisition.ClientID)
		assert.Nil(acquisition.AgentID)
		assert.Equal(enums.PaymentSchemeInstallment, acquisition.PaymentScheme)
		assert.Equal(enums.PaymentPeriodMonthly, acquisition.PaymentPeriod)
		assert.Equal(12, acquisition.Terms)
		assert.NotNil(acquisition.AcquiredAt)
		assert.Nil(acquisition.CompletedAt)

		// assert created invoices
		invoices := make([]*models.InvoiceModel, 0)
		cursor, err := db.Collection(enums.CollInvoices).Find(context.TODO(), bson.M{"propertyID": property2.ID})
		assert.Nil(err)
		err = cursor.All(context.TODO(), &invoices)
		assert.Nil(err)
		assert.Len(invoices, 13)
		invoice := invoices[0]
		assert.Equal(businessID, *invoice.From.ID)
		assert.Equal(userID, invoice.CreatedBy)
		assert.EqualValues(10000, invoice.Total)
		assert.Equal(fmt.Sprintf("%s - 0/12", property2.Name), invoice.Name)
		assert.EqualValues(2, invoice.No)

		// first recurring payment
		invoice = invoices[1]
		assert.EqualValues(832504, invoice.Total)
		assert.Equal(fmt.Sprintf("%s - 1/12", property2.Name), invoice.Name)
		assert.EqualValues(3, invoice.No)

		invoice = invoices[12]
		assert.EqualValues(832504, invoice.Total)
		assert.Equal(fmt.Sprintf("%s - 12/12", property2.Name), invoice.Name)
		assert.EqualValues(14, invoice.No)
		// TODO: assert agent commissions
	})

	t.Run("It should validate inputs", func(t *testing.T) {
		assert := assert.New(t)
		cases := []struct {
			err     string
			payload fiber.Map
		}{
			{
				errors.ErrNotFoundProperty,
				fiber.Map{
					"i": 1,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"propertyID": "someid",
					"i":          2,
				},
			},
			{
				errors.ErrNotFoundProperty,
				fiber.Map{
					"propertyID": primitive.NewObjectID(),
					"i":          3,
				},
			},
			{
				errors.ErrPropertyAlreadyTaken,
				fiber.Map{
					"propertyID": property2.ID,
					"i":          4,
				},
			},
			{
				errors.ErrNotFoundPerson,
				fiber.Map{
					"propertyID": property3.ID,
					"i":          5,
				},
			},
			{
				errors.ErrNotFoundPerson,
				fiber.Map{
					"propertyID": property3.ID,
					"clientID":   primitive.NewObjectID(),
					"i":          6,
				},
			},
			{
				errors.ErrNotFoundPerson,
				fiber.Map{
					"propertyID": property3.ID,
					"clientID":   person1.ID,
					"agentID":    primitive.NewObjectID(),
					"i":          7,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"propertyID":    property3.ID,
					"clientID":      person1.ID,
					"agentID":       person2.ID,
					"paymentScheme": "notcashorinstallment",
					"i":             8,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"propertyID":    property3.ID,
					"clientID":      person1.ID,
					"agentID":       person2.ID,
					"paymentScheme": enums.PaymentSchemeInstallment,
					"paymentPeriod": "notmonthlyoryearly",
					"i":             9,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"propertyID":    property3.ID,
					"clientID":      person1.ID,
					"agentID":       person2.ID,
					"paymentScheme": enums.PaymentSchemeInstallment,
					"paymentPeriod": enums.PaymentPeriodMonthly,
					"terms":         0,
					"i":             10,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"propertyID":    property3.ID,
					"clientID":      person1.ID,
					"agentID":       person2.ID,
					"paymentScheme": enums.PaymentSchemeInstallment,
					"paymentPeriod": enums.PaymentPeriodMonthly,
					"terms":         -3,
					"i":             11,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"propertyID":    property3.ID,
					"clientID":      person1.ID,
					"agentID":       person2.ID,
					"paymentScheme": enums.PaymentSchemeInstallment,
					"paymentPeriod": enums.PaymentPeriodMonthly,
					"terms":         12,
					"downPayment":   0,
					"i":             12,
				},
			},
			{
				errors.ErrInputInvalid,
				fiber.Map{
					"propertyID":    property3.ID,
					"clientID":      person1.ID,
					"agentID":       person2.ID,
					"paymentScheme": enums.PaymentSchemeInstallment,
					"paymentPeriod": enums.PaymentPeriodMonthly,
					"terms":         12,
					"downPayment":   -100,
					"i":             13,
				},
			},
		}

		for _, item := range cases {
			req := helpers.DoRequest("POST", path, item.payload, token1)
			res, err := app.Test(req, -1)

			// assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, item.payload)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(item.err), response.Message, item.payload)
		}
	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not acquire from another business", func(t *testing.T) {
			assert := assert.New(t)
			data := fiber.Map{
				"propertyID":    property4.ID,
				"clientID":      person1.ID,
				"paymentScheme": enums.PaymentSchemeCash,
				"agentID":       person2.ID,
			}
			req := helpers.DoRequest("POST", path, data, token1)
			res, err := app.Test(req, -1)
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, res)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(errors.ErrNotFoundProperty), response.Message, response)

		})
	})
}
