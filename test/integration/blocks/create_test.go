package blocks

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/test/helpers"
)

func TestCreateBlock(t *testing.T) {
	log.Println("-- test create block --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/blocks"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	invoice1 := helpers.InvoiceFixture(app, token1, nil, 0)
	userID, businessID, _ := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should create a block", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"type":       enums.InvoiceBlockIntro,
			"entityType": enums.EntityInvoice,
			"entityID":   invoice1.ID,
		}
		req := helpers.DoRequest("POST", path, data, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		ress, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		// decode response
		m, err := json.Marshal(ress)
		assert.Nil(err)
		response := &models.BlockModel{}
		err = json.Unmarshal(m, &response)
		assert.Nil(err)
		assert.False(response.ID.IsZero())
		assert.Equal(enums.EntityInvoice, response.EntityType)
		assert.Equal(enums.InvoiceBlockIntro, response.Type)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.CreatedBy)
		assert.Equal(userID, response.UpdatedBy)
	})

	t.Run("It should create an invoice item block", func(t *testing.T) {
		assert := assert.New(t)
		fakeTitle := "Item A"
		fakeDescription := "Description A"
		fakeAmount := 10000
		fakeQuantity := 2
		fakeTax := "10%"
		fakeDiscount := "5%"
		data := fiber.Map{
			"type":        enums.InvoiceBlockItem,
			"entityType":  enums.EntityInvoice,
			"entityID":    invoice1.ID,
			"title":       fakeTitle,
			"description": fakeDescription,
			"amount":      fakeAmount,
			"quantity":    fakeQuantity,
			"tax":         fakeTax,
			"discount":    fakeDiscount,
			"addOrLess": []fiber.Map{
				{
					"name":     "tax",
					"value":    fakeTax,
					"less":     false,
					"fromBase": true,
				},
				{
					"name":     "tax",
					"value":    fakeDiscount,
					"less":     true,
					"fromBase": false,
				},
			},
		}
		req := helpers.DoRequest("POST", path, data, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		ress, err := helpers.GetResponseMap(res)
		assert.Nil(err)
		// decode response
		m, err := json.Marshal(ress)
		assert.Nil(err)
		response := &models.InvoiceItemBlockModel{}
		err = json.Unmarshal(m, &response)
		assert.Nil(err)
		assert.False(response.ID.IsZero())
		assert.Equal(enums.EntityInvoice, response.EntityType)
		assert.Equal(enums.InvoiceBlockItem, response.Type)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.CreatedBy)
		assert.Equal(userID, response.UpdatedBy)
		assert.Equal(fakeTitle, response.Title)
		assert.Equal(fakeDescription, response.Description)
		assert.EqualValues(fakeAmount, response.Amount)
		assert.EqualValues(fakeQuantity, response.Quantity)
		als := response.AddOrLess
		// alsid := als[0].ID
		assert.False(als[0].ID.IsZero())
		assert.Equal("tax", als[0].Name)
		assert.Equal(fakeTax, als[0].Value)
		assert.False(als[0].Less)
		assert.True(als[0].FromBase)
		assert.Equal(fakeDiscount, als[1].Value)
		assert.True(als[1].Less)
		// assert.EqualValues(fakeTax, response.Tax)
		// assert.Equal(fakeDiscount, response.Discount)
		// assert.EqualValues(2000, response.TaxAmount)
		// assert.EqualValues(1100, response.DiscountAmount)
		assert.EqualValues(20900, response.Total)
	})

}
