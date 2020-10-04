package files_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/mewben/arrstate/pkg/api/files"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/test/helpers"
)

var testData1 = fiber.Map{
	"title":    "Test File 1",
	"ext":      "jpg",
	"size":     12353,
	"mimeType": "image/*",
}
var testData2 = fiber.Map{
	"title":    "Test File 2",
	"ext":      "gif",
	"size":     22343,
	"mimeType": "image/*",
}

// CreateFixture -
func CreateFixture(app *fiber.App, token string, payload interface{}) *models.FileModel {
	req := helpers.DoRequest("POST", "/api/files", payload, token)
	res, err := app.Test(req, -1)
	if err != nil {
		log.Fatalln("err fixture create file", err)
	}

	response, err := GetResponse(res)
	if err != nil {
		log.Fatalln("err fixture create file response", err)
	}

	return response
}

// GetResponse -
func GetResponse(res *http.Response) (*models.FileModel, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &models.FileModel{}
	err = json.Unmarshal(body, &response)
	return response, err
}

// GetResponses -
func GetResponses(res *http.Response) (*files.ResponseList, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &files.ResponseList{}
	err = json.Unmarshal(body, &response)
	return response, err
}
