package blocks

import (
	"log"
	"os"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	log.Println("-- test create block --")
	os.Setenv("ENV", "TESTING")
	// db := startup.Init()
	// app := pkg.SetupBackend(db)
	// path := "/api/blocks"

	// // setup
	// helpers.CleanupFixture(db)
	// token1 := helpers.SignupFixture(app, 0)
	// invoice1 := helpers.InvoiceFixture(app, token1, 0)
	// userID, businessID := helpers.CheckJWT(token1, assert.New(t))

	// t.Run("It should create a block", func(t *testing.T) {
	// 	assert := assert.New(t)
	// 	query := url.Values{}
	// 	query.Add("entityType", enums.EntityInvoice)
	// 	data := fiber.Map{
	// 		"type": enums.InvoiceBlockIntro,

	// 	}
	// })

}
