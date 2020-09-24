package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mewben/arrstate/pkg/models"
)

var FakeProject [2]fiber.Map

func init() {
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
