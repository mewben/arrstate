package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mewben/arrstate/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var FakeFile [2]fiber.Map

func init() {
	// files
	address := models.NewAddressModel()
	address.Country = "PH"
	address.State = "Bohol"
	FakeFile[0] = fiber.Map{
		"title":      "testFile",
		"ext":        "jpg",
		"mimeType":   "image/*",
		"size":       12353,
		"url":        "http://url.com",
		"entityType": "invoice",
		"entityID":   primitive.NewObjectID(),
		"type":       "file",
		"link":       "",
	}
	FakeFile[1] = fiber.Map{
		"title": "testFile2",
	}
}

// FileFixture -
func FileFixture(app *fiber.App, token string, n int) *models.FileModel {

	req := DoRequest("POST", "/api/files", FakeFile[n], token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err app test project", err)
	}

	response, err := GetResponse(res, "file")
	if err != nil {
		log.Fatalln("err get response file", err)
	}

	return response.(*models.FileModel)
}
