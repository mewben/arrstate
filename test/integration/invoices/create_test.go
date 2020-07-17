package invoices

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/test/helpers"
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

		blocks := []fiber.Map{
			{
				"type": enums.InvoiceBlockIntro,
			},
			{
				"type":        enums.InvoiceBlockItem,
				"title":       "Item A",
				"description": "Description A",
				"amount":      100,
				"quantity":    2,
				"tax":         10, // 10%,
				"discount":    "5%",
			},
			{
				"type":        enums.InvoiceBlockItem,
				"title":       "Item B",
				"description": "Description B",
				"amount":      200,
				"quantity":    1,
			},
		}

		data := fiber.Map{
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
			"blocks":     blocks,
			"discount":   "10%",
			"tax":        20, // 20%
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
		assert.Equal(person1.ID, response.From.ID)
		assert.Equal(enums.EntityPerson, response.From.EntityType)
		assert.Equal(person2.ID, response.To.ID)
		assert.Equal(enums.EntityPerson, response.To.EntityType)
		assert.Equal(project1.ID, response.PropertyID)
		assert.Equal(property1.ID, response.PropertyID)
		assert.Equal(issueDate, response.IssueDate)
		assert.Equal(dueDate, response.DueDate)
		assert.Len(response.Blocks, 3)
		assert.Equal(enums.StatusPending, response.Status)
		assert.Equal("1", response.InvoiceSeq)
		assert.EqualValues(20, response.Tax)
		assert.EqualValues(81.8, response.TaxAmount)
		assert.Equal("10%", response.Discount)
		assert.EqualValues(49.08, response.DiscountAmount)
		assert.EqualValues(409, response.SubTotal)
		assert.EqualValues(441.72, response.Total)

		// assert blocks
		filter := bson.D{
			{
				Key: "_id",
				Value: bson.D{
					{
						Key:   "$in",
						Value: response.Blocks,
					},
				},
			},
		}
		blocksCursor, err := db.Collection(enums.CollBlocks).Find(context.TODO(), filter)
		assert.Nil(err)
		blocksList := make([]*models.BlockI, 0)
		err = blocksCursor.All(context.TODO(), &blocksList)
		assert.Nil(err)
		log.Println("blocksList", blocksList)
	})

	t.Run("It should create invoice with provided blocks", func(t *testing.T) {

	})

	t.Run("It should validate input", func(t *testing.T) {

	})
}
