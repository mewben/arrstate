package blocks

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
	"github.com/mewben/arrstate/pkg/api/blocks"
	"github.com/mewben/arrstate/test/helpers"
)

func TestGetBlocks(t *testing.T) {
	log.Println("-- test get blocks --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/blocks/get"

	// setup
	helpers.CleanupFixture(db)
	token1 := helpers.SignupFixture(app, 0)
	invoice1 := helpers.InvoiceFixture(app, token1, nil, 1)

	t.Run("It should get blocks by ids", func(t *testing.T) {
		assert := assert.New(t)
		data := fiber.Map{
			"ids":        invoice1.Blocks,
			"entityType": enums.EntityInvoice,
			"entityID":   invoice1.ID,
		}
		req := helpers.DoRequest("POST", path, data, token1)
		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(200, res.StatusCode, res)
		ress, err := helpers.GetResponse(res, "blocks")
		assert.Nil(err)
		response := ress.(*blocks.ResponseList)
		assert.Len(response.Data, 3)
		assert.Equal(response.Total, 3)

	})

}
