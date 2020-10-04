package invoices

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services"
	"github.com/mewben/arrstate/test/helpers"
)

func TestPayInvoice(t *testing.T) {
	log.Println("-- test pay invoice --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/invoices/pay"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	invoice1 := helpers.InvoiceFixture(app, token1, nil, 0)
	invoice2 := helpers.InvoiceFixture(app, token1, nil, 0)
	// userID, businessID := helpers.CheckJWT(token1, assert.New(t))

	t.Run("It should pay an invoice", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"receiptNo": "1234",
			"invoiceID": invoice1.ID,
		}
		log.Println("dataaa", data)
		req := helpers.DoRequest("POST", path, data, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "invoice")
		assert.Nil(err)
		response := ress.(*models.InvoiceModel)
		assert.Equal(enums.StatusPaid, response.Status)
		assert.Equal("1234", response.ReceiptNo)
		assert.NotNil(response.PaidAt)
	})

	t.Run("It should validate inputs", func(t *testing.T) {
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
					"receiptNo": "1234",
					"invoiceID": "someid",
				},
			},
			{
				errors.ErrNotFoundInvoice,
				fiber.Map{
					"receiptNo": "12345",
					"invoiceID": primitive.NewObjectID(),
				},
			},
			{
				errors.ErrInvoiceAlreadyPaid,
				fiber.Map{
					"receiptNo": "djf34",
					"invoiceID": invoice1.ID,
				},
			},
			{
				errors.ErrDuplicateReceipt,
				fiber.Map{
					"receiptNo": "1234",
					"invoiceID": invoice2.ID,
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
}
