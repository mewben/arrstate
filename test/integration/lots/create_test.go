package lots

import (
	"log"
	"os"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/test/helpers"
)

func TestCreateLot(t *testing.T) {
	log.Println("-- test create lot --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/lots"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app, 1)
	project := helpers.ProjectFixture(app, authResponse.Token, 1)

	t.Run("It should create lot", func(t *testing.T) {
		assert := assert.New(t)
		fakeProjectID := project.ID.Hex()
		fakeName := "Fake lot"
		fakeArea := 53.5
		fakePrice := 100.5
		fakePriceAddon := 10.1
		fakeNotes := "Sample Notes"
		fakeImages := []*models.ImageModel{
			{
				ID:          primitive.NewObjectID(),
				Src:         "src",
				Alt:         "alt",
				Description: "description",
			},
		}

		data := fiber.Map{
			"projectID":  fakeProjectID,
			"name":       fakeName,
			"area":       fakeArea,
			"price":      fakePrice,
			"priceAddon": fakePriceAddon,
			"notes":      fakeNotes,
			"images":     fakeImages,
		}
		req := helpers.DoRequest("POST", path, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		response, err := helpers.GetResponseLot(res)
		assert.Nil(err)
		assert.Equal(authResponse.CurrentBusiness.ID, response.BusinessID)
		assert.Equal(authResponse.CurrentUser.User.ID, response.CreatedBy)
		assert.NotEmpty(response.ID)
		assert.Equal(fakeProjectID, response.ProjectID.Hex())
		assert.Equal(fakeName, response.Name)
		assert.EqualValues(fakeArea, response.Area)
		assert.EqualValues(fakePrice, response.Price)
		assert.EqualValues(fakePriceAddon, response.PriceAddon)
		assert.Equal(fakeNotes, response.Notes)
		assert.Len(response.Images, 1)
		assert.Equal(fakeImages[0].ID, response.Images[0].ID)
		assert.Equal(fakeImages[0].Src, response.Images[0].Src)
		assert.Equal(fakeImages[0].Alt, response.Images[0].Alt)
		assert.Equal(fakeImages[0].Description, response.Images[0].Description)
	})

	t.Run("It should validate lot inputs", func(t *testing.T) {

	})

	t.Run("It should throw if project not found", func(t *testing.T) {

	})

}