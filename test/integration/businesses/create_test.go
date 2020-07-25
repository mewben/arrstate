package businesses

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateBusiness(t *testing.T) {
	log.Println("-- BUSINESS.CREATE --")

	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/businesses"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)

	t.Run("It should create business", func(t *testing.T) {
		assert := assert.New(t)
		fakeName := "Test Domain"
		fakeDomain := "some-domain"
		data := fiber.Map{
			"name":   fakeName,
			"domain": fakeDomain,
		}
		req := helpers.DoRequest("POST", path, data, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "business")
		assert.Nil(err)
		response := ress.(*models.BusinessModel)
		assert.False(response.ID.IsZero())
		assert.Equal(fakeName, response.Name)
		assert.Equal(fakeDomain, response.Domain)
		assert.EqualValues(1, response.Invoices.NextSeq)
	})

	t.Run("It should set default values", func(t *testing.T) {

	})

	t.Run("It should validate input", func(t *testing.T) {
		// include invalid domain characters

	})

	t.Run("It should throw if domain is already taken", func(t *testing.T) {

	})

}
