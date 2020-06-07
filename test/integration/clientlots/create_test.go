package clientlots

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
	"github.com/mewben/realty278/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateClientLot(t *testing.T) {
	log.Println("-- test create clientlot --")
	os.Setenv("ENV", "TESTING")
	db := startup.Init()
	app := pkg.SetupBackend(db)
	path := "/api/clientlots"

	// setup
	helpers.CleanupFixture(db)
	_, authResponse := helpers.SignupFixture(app, 1)
	project1 := helpers.ProjectFixture(app, authResponse.Token, 1)
	lot1 := helpers.LotFixture(app, authResponse.Token, project1.ID, 1)
	person1 := helpers.PersonFixture(app, authResponse.Token, 1)
	person2 := helpers.PersonFixture(app, authResponse.Token, 2)

	t.Run("It should attach a person to a lot", func(t *testing.T) {
		assert := assert.New(t)
		fakeDownPayment := 12.5
		fakeMonths := 12
		fakeMonthly := 4.5
		fakeDate := time.Now().Add(time.Hour * 24)
		data := fiber.Map{
			"lotID":       lot1.ID,
			"clientID":    person1.ID,
			"agentID":     person2.ID,
			"downPayment": fakeDownPayment,
			"months":      fakeMonths,
			"monthly":     fakeMonthly,
			"date":        fakeDate,
		}
		req := helpers.DoRequest("POST", path, data, authResponse.Token)

		res, err := app.Test(req, -1)
		assert.Nil(err)
		assert.Equal(201, res.StatusCode, res)
		resp, err := helpers.GetResponseClientLot(res)
		assert.Nil(err)
		response := resp.ClientLot
		updatedLot := resp.Lot
		assert.Equal(authResponse.CurrentBusiness.ID, response.BusinessID)
		assert.Equal(authResponse.CurrentUser.User.ID, response.CreatedBy)
		assert.False(response.ID.IsZero())
		assert.Equal(lot1.ID, response.LotID)
		assert.Equal(person1.ID, response.ClientID)
		assert.Equal(person2.ID, *response.AgentID)
		assert.Equal(enums.StatusPending, response.Status)
		assert.EqualValues(lot1.Price+lot1.PriceAddon, response.Price)
		assert.EqualValues(fakeDownPayment, response.DownPayment)
		assert.Equal(fakeMonths, response.Months)
		assert.EqualValues(fakeMonthly, response.Monthly)
		log.Println("-fakeDateResponseDate", fakeDate, fakeDate.Sub(response.Date).Hours(), "sss")
		log.Println("responseDate", response.Date)
		assert.EqualValues(0, int64(fakeDate.Sub(response.Date).Seconds()))

		// the lot must have clientLotID and status=pending
		assert.Equal(enums.StatusPending, updatedLot.Status)
		assert.Equal(response.ID, *updatedLot.ClientLotID)
	})

	// TODO: move this to edit
	t.Run("It should generate invoices after attach", func(t *testing.T) {

	})

	t.Run("It should validate inputs", func(t *testing.T) {

	})

	t.Run("Permissions", func(t *testing.T) {
		// TODO
	})

}
