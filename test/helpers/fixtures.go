package helpers

import (
	"context"
	"log"
	"math/rand"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/auth"
	"github.com/mewben/realty278/pkg/models"
)

// fixture variables
var signup1 *auth.SignupPayload
var signup2 *auth.SignupPayload
var project1 fiber.Map
var project2 fiber.Map
var lot1 fiber.Map
var lot2 fiber.Map
var person1 fiber.Map
var person2 fiber.Map

func init() {

	// business
	signup1 = &auth.SignupPayload{
		GivenName:  "testgn",
		FamilyName: "testfn",
		Business:   "Test Business",
		Domain:     "test-domain",
		Email:      "test@email.com",
		Password:   "password",
	}
	signup2 = &auth.SignupPayload{
		GivenName:  "testgn2",
		FamilyName: "testfn2",
		Business:   "Test Business2",
		Domain:     "test-domain2",
		Email:      "test2@email.com",
		Password:   "password2",
	}

	// people
	person1 = fiber.Map{
		"email":     "email1@email.com",
		"role":      enums.RoleAgent,
		"givenName": "given",
	}
	person2 = fiber.Map{
		"email":     "email2@email.com",
		"role":      enums.RoleAgent,
		"givenName": "given2",
	}

	// projects
	address := models.NewAddressModel()
	address.Country = "PH"
	address.State = "Bohol"
	project1 = fiber.Map{
		"name":    "testproject",
		"address": address,
		"area":    100.5,
		"notes":   "Sample Notes",
	}
	project2 = fiber.Map{
		"name": "testproject2",
	}

	// lots
	lot1 = fiber.Map{
		"name":       "testlot",
		"area":       1.5,
		"price":      100.5,
		"priceAddon": 114,
		"notes":      "Sample Notes",
	}
	lot2 = fiber.Map{
		"name":  "testlot2",
		"area":  2.5,
		"price": 12.3,
	}

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
		enums.CollLots,
	}
	for _, col := range collections {
		db.Collection(col).DeleteMany(context.TODO(), bson.D{})
	}
}

// SignupFixture -
func SignupFixture(app *fiber.App, n int) (*auth.SignupPayload, *models.AuthSuccessResponse) {
	payload := signup1
	if n == 2 {
		payload = signup2
	}
	req := DoRequest("POST", "/auth/signup", payload, "")
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test signup", err)
	}

	log.Println("resauth 1", res)

	response, err := GetResponseAuth(res)
	if err != nil {
		log.Fatalln("err get response auth", err)
	}

	return payload, response
}

// ProjectFixture -
func ProjectFixture(app *fiber.App, token string, n int) *models.ProjectModel {
	payload := project1
	if n == 2 {
		payload = project2
	}

	req := DoRequest("POST", "/api/projects", payload, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test project", err)
	}

	response, err := GetResponseProject(res)
	if err != nil {
		log.Fatalln("err get response project", err)
	}

	return response
}

// LotFixture -
func LotFixture(app *fiber.App, token string, projectID primitive.ObjectID, n int) *models.LotModel {
	payload := lot1
	if n == 2 {
		payload = lot2
	}

	payload["projectID"] = projectID

	req := DoRequest("POST", "/api/lots", payload, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test lot", err)
	}

	response, err := GetResponseLot(res)
	if err != nil {
		log.Fatalln("err get response lot", err)
	}

	return response
}

// PersonFixture -
func PersonFixture(app *fiber.App, token string, n int) *models.PersonModel {
	payload := person1
	if n == 2 {
		payload = person2
	}

	req := DoRequest("POST", "/api/people", payload, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test people", err)
	}

	response, err := GetResponsePerson(res)
	if err != nil {
		log.Fatalln("err get response people", err)
	}

	return response
}
