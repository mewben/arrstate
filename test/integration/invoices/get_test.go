package invoices

import (
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/api/invoices"
	"github.com/mewben/realty278/test/helpers"
)

func TestGetInvoices(t *testing.T) {
	log.Println("-- test get invoices --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/invoices"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	property1 := helpers.AcquireFixture(app, token1)

	t.Run("It should get the list of invoices in a property", func(t *testing.T) {
		assert := assert.New(t)
		query := url.Values{}
		query.Add("propertyID", property1.ID.Hex())
		req := helpers.DoRequest("GET", path+"?"+query.Encode(), nil, token1)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "invoices")
		assert.Nil(err)
		response := ress.(*invoices.ResponseList)
		assert.Len(response.Data, 13)
		assert.Equal(response.Total, 13)
	})

}
