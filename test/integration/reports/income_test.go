package reports

import (
	"context"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/api/reports"
	"github.com/mewben/arrstate/test/helpers"
)

func TestIncomeReports(t *testing.T) {
	log.Println("-- test get income reports --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/reports"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	token2 := helpers.SignupFixture(app, 1)
	invoice1 := helpers.InvoiceFixture(app, token1, nil, 0)
	invoice2 := helpers.InvoiceFixture(app, token1, nil, 0)
	invoice3 := helpers.InvoiceFixture(app, token1, nil, 0)
	invoice4 := helpers.InvoiceFixture(app, token2, nil, 0)

	t.Run("It should return income reports with provided date range", func(t *testing.T) {
		assert := assert.New(t)
		helpers.ReceiptFixture(app, token1, fiber.Map{
			"receiptNo": "1234",
			"invoiceID": invoice1.ID,
		})
		helpers.ReceiptFixture(app, token1, fiber.Map{
			"receiptNo": "1235",
			"invoiceID": invoice2.ID,
		})
		helpers.ReceiptFixture(app, token1, fiber.Map{
			"receiptNo": "1236",
			"invoiceID": invoice3.ID,
		})
		helpers.ReceiptFixture(app, token2, fiber.Map{
			"receiptNo": "1237",
			"invoiceID": invoice4.ID,
		})

		// manually change the paidAt of invoiceR2
		paidAt := time.Now().Add(-40 * 24 * 60 * 60 * time.Second) // - 40 days
		upd := bson.D{
			{
				Key:   "$set",
				Value: bson.M{"paidAt": paidAt},
			},
		}
		db.Collection(enums.CollInvoices).FindOneAndUpdate(context.TODO(), bson.M{"_id": invoice2.ID}, upd)

		// prepare query
		now := time.Now()
		format := "01-02-2006"
		query := url.Values{}
		query.Add("from", now.Add(-30*24*60*60*time.Second).Format(format))
		query.Add("to", now.Format(format))
		req := helpers.DoRequest("GET", path+"/income?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "reports")
		assert.Nil(err)
		response := ress.(*reports.ResponseList)
		assert.Equal(response.Total, 2)

	})
}
