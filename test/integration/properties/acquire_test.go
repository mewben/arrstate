package properties

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/test/helpers"
)

func TestAcquireProperty(t *testing.T) {
	log.Println("-- test create clientproperty --")
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
	property3 := helpers.PropertyFixture(app, token2, &project2.ID, 0)
	person1 := helpers.PersonFixture(app, token1, 0)
	person2 := helpers.PersonFixture(app, token1, 1)
	person3 := helpers.PersonFixture(app, token2, 1)
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
	})

	t.Run("It should let a client acquire a property in installment", func(t *testing.T) {
		// assert := assert.New(t)
		data := fiber.Map{
			"propertyID":    property2.ID,
			"clientID":      person1.ID,
			"paymentScheme": enums.PaymentSchemeCash,
		}

		log.Println("data", data)

	})

	t.Run("Permissions", func(t *testing.T) {
		t.Run("It should not acquire from another business", func(t *testing.T) {
			// assert := assert.New(t)
			log.Println("property3", property3)
			log.Println("person3", person3)
			log.Println("userID", userID)
			log.Println("businessID", businessID)
		})
	})
}
