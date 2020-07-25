package invoices

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/test/helpers"
)

func TestCreateInvoice(t *testing.T) {
	log.Println("-- test create invoice --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/invoices"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	person1 := helpers.PersonFixture(app, token1, 0)
	person2 := helpers.PersonFixture(app, token1, 1)
	project1 := helpers.ProjectFixture(app, token1, 0)
	property1 := helpers.PropertyFixture(app, token1, &project1.ID, 0)
	userID, businessID := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should create invoice", func(t *testing.T) {
		assert := assert.New(t)
		issueDate := time.Now()
		dueDate := time.Now().Add(time.Hour * 24)

		blocksData := []fiber.Map{
			{
				"type":        enums.InvoiceBlockItem,
				"title":       "Item A",
				"description": "Description A",
				"amount":      10000,
				"quantity":    2,
				"tax":         1000, // 10%,
				"discount":    "5%",
			},
			{
				"type":        enums.InvoiceBlockItem,
				"title":       "Item B",
				"description": "Description B",
				"amount":      20000,
				"quantity":    1,
			},
		}

		data := fiber.Map{
			"name": "Invoice 1",
			"from": fiber.Map{
				"_id":        person1.ID,
				"entityType": enums.EntityPerson,
			},
			"to": fiber.Map{
				"_id":        person2.ID,
				"entityType": enums.EntityPerson,
			},
			"projectID":  project1.ID,
			"propertyID": property1.ID,
			"issueDate":  issueDate,
			"dueDate":    dueDate,
			"blocks":     blocksData,
			"discount":   "10%",
			"tax":        2000, // 20%
		}

		// computations block
		// item | amount 	| qty | tax | discount 	| total
		// A		| 100			| 2		| 10%	| 5%				| 209
		// B		| 200			| 1		| 0		| 0					| 200
		// 100*2 = 200
		// 200*.10 = 20 // tax
		// 200+20 = 220
		// 220 * 0.05 = 11 // discount
		// total = 220 - 11 = 209

		// Using int
		// 10000*2 = 20,000
		// 20,000*1000/10000 = 2,000 // tax
		// 20,000+2,000 = 22,000
		// 22,000 * 500/10000 = 1,100 // discount
		// total = 22,000 - 1,100 = 20,900

		// computations summary
		// subTotal = 409
		// tax = 20%, discount = 10%
		// 409*.20 = 81.8 // tax
		// 409+81.8 = 490.8
		// 490.8*.1 = 49.08 // discount
		// total = 409+81.8-49.08 = 441.72

		req := helpers.DoRequest("POST", path, data, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "invoice")
		assert.Nil(err)
		response := ress.(*models.InvoiceModel)
		assert.Equal(businessID, response.BusinessID)
		assert.Equal(userID, response.CreatedBy)
		assert.Equal(userID, response.UpdatedBy)
		assert.False(response.ID.IsZero())
		assert.Equal(enums.StatusDraft, response.Status)
		assert.Equal(person1.ID, *response.From.ID)
		assert.Equal(enums.EntityPerson, response.From.EntityType)
		assert.Equal(person2.ID, *response.To.ID)
		assert.Equal(enums.EntityPerson, response.To.EntityType)
		assert.Equal(project1.ID, *response.ProjectID)
		assert.Equal(property1.ID, *response.PropertyID)
		isd1, _ := time.Parse("02 Jan 06 15:04", issueDate.String())
		isd2, _ := time.Parse("02 Jan 06 15:04", response.IssueDate.String())
		assert.Equal(isd1, isd2)
		isd1, _ = time.Parse("02 Jan 06 15:04", dueDate.String())
		isd2, _ = time.Parse("02 Jan 06 15:04", response.DueDate.String())
		assert.Equal(isd1, isd2)
		assert.Len(response.Blocks, 4)
		assert.Equal(enums.StatusDraft, response.Status)
		assert.Equal("Invoice 1", response.Name)
		assert.EqualValues(1, response.No)
		assert.EqualValues(2000, response.Tax)
		assert.EqualValues(8180, response.TaxAmount)
		assert.Equal("10%", response.Discount)
		assert.EqualValues(4908, response.DiscountAmount)
		assert.EqualValues(40900, response.SubTotal)
		assert.EqualValues(44172, response.Total)
	})

	t.Run("It should create invoice with provided blocks", func(t *testing.T) {

	})

	t.Run("It should validate input", func(t *testing.T) {

	})
}
