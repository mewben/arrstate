package helpers

import (
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/auth"
	"github.com/mewben/arrstate/pkg/models"
)

// fixture variables - array not slice so we can directly set index

// FakeSignup -
var (
	FakeSignup [2]*auth.SignupPayload
	// FakeProject  [2]fiber.Map
	FakeProperty [2]fiber.Map
	FakePerson   [3]fiber.Map
	FakeInvoice  [2]fiber.Map
)

func init() {
	// business
	FakeSignup[0] = &auth.SignupPayload{
		Name: models.PersonName{
			First: "first",
			Last:  "last",
		},
		Business: "Test Business",
		Domain:   "test-domain",
		Email:    "test@email.com",
		Password: "password",
	}
	FakeSignup[1] = &auth.SignupPayload{
		Name: models.PersonName{
			First: "first2",
			Last:  "last2",
		},
		Business: "Test Business2",
		Domain:   "test-domain2",
		Email:    "test2@email.com",
		Password: "password2",
	}

	// people
	FakePerson[0] = fiber.Map{
		"email": "email1@email.com",
		"role":  []string{enums.RoleAgent},
		"name": fiber.Map{
			"first": "first",
		},
	}
	FakePerson[1] = fiber.Map{
		"email": "email2@email.com",
		"role":  []string{enums.RoleClient, enums.RoleAgent},
		"name": fiber.Map{
			"first": "first2",
		},
		"commissionPerc": 1000,
	}
	FakePerson[2] = fiber.Map{
		"email": "email3@email.com",
		"role":  []string{enums.RoleClient, enums.RoleAgent},
		"name": fiber.Map{
			"first": "first3",
		},
		"commissionPerc": 500,
	}

	// properties
	FakeProperty[0] = fiber.Map{
		"name":       "testproperty",
		"type":       enums.PropertyTypeLot,
		"area":       1.5,
		"price":      10000050,
		"priceAddon": 11400,
		"notes":      "Sample Notes",
	}
	FakeProperty[1] = fiber.Map{
		"name":  "testproperty2",
		"type":  enums.PropertyTypeHouse,
		"area":  2.5,
		"price": 12000030,
	}

	// invoices
	FakeInvoice[0] = fiber.Map{
		// "tax":       1250,
		// "discount":  "5%",
		"issueDate": time.Now(),
		"dueDate":   time.Now().Add(24 * time.Hour),
		"blocks": []fiber.Map{
			{
				"type":   enums.InvoiceBlockItem,
				"amount": 100000,
			},
		},
		"addOrLess": []fiber.Map{
			{
				"name":     "tax",
				"value":    "12.5%",
				"less":     false,
				"fromBase": true,
			},
			{
				"name":     "discount",
				"value":    "5%",
				"less":     true,
				"fromBase": true,
			},
		},
	}
	FakeInvoice[1] = fiber.Map{}

}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// SignupFixture -
func SignupFixture(app *fiber.App, n int) string {
	req := DoRequest("POST", "/auth/signup", FakeSignup[n], "")
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test signup", err)
	}

	response, err := GetResponseMap(res)
	if err != nil {
		log.Fatalln("err get response auth", err)
	}

	signinPayload := fiber.Map{
		"grant_type": "device_code",
		"deviceCode": response["deviceCode"],
	}
	req = DoRequest("POST", "/auth/signin", signinPayload, "")
	req.Header.Add("origin", "http://"+FakeSignup[n].Domain+".example.com")
	res, err = app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test signup", err)
	}

	response, err = GetResponseMap(res)
	if err != nil {
		log.Fatalln("err get response auth", err)
	}

	return response["token"].(string)
}

// PropertyFixture -
func PropertyFixture(app *fiber.App, token string, projectID *primitive.ObjectID, n int) *models.PropertyModel {
	payload := FakeProperty[n]
	if projectID != nil {
		payload["projectID"] = projectID
	} else {
		delete(payload, "projectID")
	}

	req := DoRequest("POST", "/api/properties", payload, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test property", err)
	}

	response, err := GetResponse(res, "property")
	if err != nil {
		log.Fatalln("err get response property", err)
	}

	return response.(*models.PropertyModel)
}

// AcquireFixture -
func AcquireFixture(app *fiber.App, token string) *models.PropertyModel {
	property := PropertyFixture(app, token, nil, 0)
	client := PersonFixture(app, token, 0)
	data := fiber.Map{
		"propertyID":    property.ID,
		"clientID":      client.ID,
		"paymentScheme": enums.PaymentSchemeInstallment,
		"paymentPeriod": enums.PaymentPeriodMonthly,
		"terms":         12,
		"downPayment":   10000,
	}
	req := DoRequest("POST", "/api/properties/acquire", data, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test acquire property", err)
	}

	response, err := GetResponse(res, "property")
	if err != nil {
		log.Fatalln("err get response acquire property", err)
	}

	return response.(*models.PropertyModel)
}

// PersonFixture -
func PersonFixture(app *fiber.App, token string, n int) *models.PersonModel {
	req := DoRequest("POST", "/api/people", FakePerson[n], token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test people", err)
	}

	response, err := GetResponse(res, "person")
	if err != nil {
		log.Fatalln("err get response people", err)
	}

	return response.(*models.PersonModel)
}

// InvoiceFixture -
func InvoiceFixture(app *fiber.App, token string, propertyID *primitive.ObjectID, n int) *models.InvoiceModel {
	payload := FakeInvoice[n]
	if propertyID != nil {
		payload["propertyID"] = propertyID
	} else {
		delete(payload, "propertyID")
	}
	req := DoRequest("POST", "/api/invoices", payload, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err invoice fixture", err)
	}

	response, err := GetResponse(res, "invoice")
	if err != nil {
		log.Fatalln("err get response invoice", err)
	}

	return response.(*models.InvoiceModel)
}

// ReceiptFixture -
func ReceiptFixture(app *fiber.App, token string, payload fiber.Map) *models.InvoiceModel {
	req := DoRequest("POST", "/api/invoices/pay", payload, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err receipt fixture", err)
	}

	response, err := GetResponse(res, "invoice")
	if err != nil {
		log.Fatalln("err get response receipt", err)
	}

	return response.(*models.InvoiceModel)
}
