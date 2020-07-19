package helpers

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/auth"
	"github.com/mewben/realty278/pkg/models"
)

// fixture variables - array not slice so we can directly set index

// FakeSignup -
var (
	FakeSignup   [2]*auth.SignupPayload
	FakeProject  [2]fiber.Map
	FakeProperty [2]fiber.Map
	FakePerson   [3]fiber.Map
	FakeInvoice  [2]fiber.Map
)

func init() {
	// business
	FakeSignup[0] = &auth.SignupPayload{
		GivenName:  "testgn",
		FamilyName: "testfn",
		Business:   "Test Business",
		Domain:     "test-domain",
		Email:      "test@email.com",
		Password:   "password",
	}
	FakeSignup[1] = &auth.SignupPayload{
		GivenName:  "testgn2",
		FamilyName: "testfn2",
		Business:   "Test Business2",
		Domain:     "test-domain2",
		Email:      "test2@email.com",
		Password:   "password2",
	}

	// people
	FakePerson[0] = fiber.Map{
		"email":     "email1@email.com",
		"role":      []string{enums.RoleAgent},
		"givenName": "given",
	}
	FakePerson[1] = fiber.Map{
		"email":     "email2@email.com",
		"role":      []string{enums.RoleClient},
		"givenName": "given2",
	}
	FakePerson[2] = fiber.Map{
		"email":     "email3@email.com",
		"role":      []string{enums.RoleClient},
		"givenName": "given3",
	}

	// projects
	address := models.NewAddressModel()
	address.Country = "PH"
	address.State = "Bohol"
	FakeProject[0] = fiber.Map{
		"name":    "testproject",
		"address": address,
		"area":    100.5,
		"notes":   "Sample Notes",
	}
	FakeProject[1] = fiber.Map{
		"name": "testproject2",
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
		"tax":       1250,
		"discount":  "5%",
		"issueDate": time.Now(),
		"dueDate":   time.Now().Add(24 * time.Hour),
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

// CleanupFixture -
func CleanupFixture(db *mongo.Database) {
	collections := []string{
		enums.CollBusinesses,
		enums.CollUsers,
		enums.CollPeople,
		enums.CollProjects,
		enums.CollProperties,
		enums.CollInvoices,
		enums.CollBlocks,
	}
	for _, col := range collections {
		db.Collection(col).DeleteMany(context.TODO(), bson.D{})
	}
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

// MeFixture -
func MeFixture(app *fiber.App, token string) *models.MeModel {
	req := DoRequest("GET", "/api/me", nil, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test me", err)
	}
	response, err := GetResponse(res, "me")
	if err != nil {
		log.Fatalln("err get response me", err)
	}

	return response.(*models.MeModel)
}

// ProjectFixture -
func ProjectFixture(app *fiber.App, token string, n int) *models.ProjectModel {

	req := DoRequest("POST", "/api/projects", FakeProject[n], token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test project", err)
	}

	response, err := GetResponse(res, "project")
	if err != nil {
		log.Fatalln("err get response project", err)
	}

	return response.(*models.ProjectModel)
}

// PropertyFixture -
func PropertyFixture(app *fiber.App, token string, projectID *primitive.ObjectID, n int) *models.PropertyModel {
	payload := FakeProperty[n]
	payload["projectID"] = projectID

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
func InvoiceFixture(app *fiber.App, token string, n int) *models.InvoiceModel {
	req := DoRequest("POST", "/api/invoices", FakeInvoice[n], token)
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
