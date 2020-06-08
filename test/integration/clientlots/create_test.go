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
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/services"
	"github.com/mewben/realty278/test/helpers"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	_, authResponse2 := helpers.SignupFixture(app, 2)
	project1 := helpers.ProjectFixture(app, authResponse.Token, 1)
	project2 := helpers.ProjectFixture(app, authResponse2.Token, 1)
	lot1 := helpers.LotFixture(app, authResponse.Token, project1.ID, 1)
	lot2 := helpers.LotFixture(app, authResponse.Token, project1.ID, 1)
	lot3 := helpers.LotFixture(app, authResponse2.Token, project2.ID, 1)
	person1 := helpers.PersonFixture(app, authResponse.Token, 1)
	person2 := helpers.PersonFixture(app, authResponse.Token, 2)
	person3 := helpers.PersonFixture(app, authResponse2.Token, 2)

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
				errors.ErrInputInvalid, // error dp
				fiber.Map{
					"x":        1,
					"lotID":    lot2.ID,
					"clientID": person1.ID,
					"agentID":  person2.ID,
					"months":   12,
					"monthly":  15.2,
					"date":     time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error dp
				fiber.Map{
					"x":           2,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": -13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error dp
				fiber.Map{
					"x":           3,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": "",
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error months
				fiber.Map{
					"x":           4,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error months
				fiber.Map{
					"x":           5,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      -1,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error months
				fiber.Map{
					"x":           6,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"months":      "",
					"downPayment": 13.4,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error monthly
				fiber.Map{
					"x":           7,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error monthly
				fiber.Map{
					"x":           8,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     -15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error monthly
				fiber.Map{
					"x":           9,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     "-15.2",
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error date
				fiber.Map{
					"x":           10,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
				},
			},
			{
				errors.ErrInputInvalid, // error date
				fiber.Map{
					"x":           11,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        "",
				},
			},
			{
				errors.ErrInputInvalid, // error date
				fiber.Map{
					"x":           12,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        "notadate",
				},
			},
			{
				errors.ErrNotFoundLot, // error lotID
				fiber.Map{
					"x":           13, // some test indicator
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error lotID
				fiber.Map{
					"x":           14,
					"lotID":       "",
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error lotID
				fiber.Map{
					"x":           15,
					"lotID":       "someid",
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrNotFoundLot, // error lotID
				fiber.Map{
					"x":           16,
					"lotID":       primitive.NewObjectID(),
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrNotFoundLot, // error lotID, lot from other business
				fiber.Map{
					"x":           17,
					"lotID":       lot3.ID,
					"clientID":    person1.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error clientID
				fiber.Map{
					"x":           18,
					"lotID":       lot2.ID,
					"clientID":    "",
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrNotFoundPerson, // error clientID
				fiber.Map{
					"x":           19,
					"lotID":       lot2.ID,
					"clientID":    primitive.NewObjectID(),
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrNotFoundPerson, // error clientID,other business
				fiber.Map{
					"x":           20,
					"lotID":       lot2.ID,
					"clientID":    person3.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrInputInvalid, // error agentID
				fiber.Map{
					"x":           21,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     "",
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrNotFoundPerson, // error agentID
				fiber.Map{
					"x":           22,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     primitive.NewObjectID(),
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrNotFoundPerson, // error agentID
				fiber.Map{
					"x":           23,
					"lotID":       lot2.ID,
					"clientID":    person1.ID,
					"agentID":     person3.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
			{
				errors.ErrLotAlreadyTaken, // error lot already has client
				fiber.Map{
					"x":           24,
					"lotID":       lot1.ID,
					"clientID":    person2.ID,
					"agentID":     person2.ID,
					"downPayment": 13.4,
					"months":      12,
					"monthly":     15.2,
					"date":        time.Now(),
				},
			},
		}

		for _, item := range cases {
			req := helpers.DoRequest("POST", path, item.payload, authResponse.Token)
			res, err := app.Test(req, -1)

			// Assert
			assert.Nil(err)
			assert.Equal(400, res.StatusCode, item.payload)
			response, err := helpers.GetResponseError(res)
			assert.Nil(err)
			assert.Equal(services.T(item.err), response.Message, item.payload)
		}

	})

	t.Run("Permissions", func(t *testing.T) {
		// TODO
	})

}
