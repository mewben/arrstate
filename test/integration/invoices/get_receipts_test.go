package invoices

import (
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/api/invoices"
	"github.com/mewben/arrstate/test/helpers"
)

func TestGetReceipts(t *testing.T) {
	log.Println("-- test get receipts --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/invoices"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	property1 := helpers.PropertyFixture(app, token1, nil, 0)
	invoice1 := helpers.InvoiceFixture(app, token1, &property1.ID, 0)
	invoice2 := helpers.InvoiceFixture(app, token1, &property1.ID, 0)
	invoice3 := helpers.InvoiceFixture(app, token1, nil, 0)
	helpers.InvoiceFixture(app, token1, nil, 0) // invoice4
	invoice1 = helpers.ReceiptFixture(app, token1, fiber.Map{
		"receiptNo": "1234",
		"invoiceID": invoice1.ID,
	})
	invoice2 = helpers.ReceiptFixture(app, token1, fiber.Map{
		"receiptNo": "1235",
		"invoiceID": invoice2.ID,
	})
	invoice3 = helpers.ReceiptFixture(app, token1, fiber.Map{
		"receiptNo": "1236",
		"invoiceID": invoice3.ID,
	})

	t.Run("It should get the list of receipts by propertyID", func(t *testing.T) {
		assert := assert.New(t)
		query := url.Values{}
		query.Add("propertyID", property1.ID.Hex())
		query.Add("status", enums.StatusPaid)
		req := helpers.DoRequest("GET", path+"?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "invoices")
		assert.Nil(err)
		response := ress.(*invoices.ResponseList)
		assert.Len(response.Data, 2)
		assert.Equal(response.Total, 2)
	})

}
