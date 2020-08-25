package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber"

	"github.com/mewben/arrstate/pkg/api/blocks"
	"github.com/mewben/arrstate/pkg/api/businesses"
	"github.com/mewben/arrstate/pkg/api/invoices"
	"github.com/mewben/arrstate/pkg/api/people"
	"github.com/mewben/arrstate/pkg/api/projects"
	"github.com/mewben/arrstate/pkg/api/properties"
	"github.com/mewben/arrstate/pkg/auth"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// GetResponseError http
func GetResponseError(res *http.Response) (*errors.HTTPError, error) {
	responseError := &errors.HTTPError{}
	if res.StatusCode == 500 {
		msg, err := GetResponseString(res)
		log.Println("- err 500:", msg)
		responseError.Message = msg
		return responseError, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return responseError, nil
	}
	err = json.Unmarshal(body, &responseError)
	return responseError, err
}

// GetResponseString -
func GetResponseString(res *http.Response) (string, error) {
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}

// GetResponseMap general error
func GetResponseMap(res *http.Response) (fiber.Map, error) {
	response := fiber.Map{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err

}

// GetResponse success
func GetResponse(res *http.Response, entity string) (interface{}, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response interface{}
	if entity == "signin" {
		response = &auth.SigninResponse{}
	} else if entity == "business" {
		response = &models.BusinessModel{}
	} else if entity == "businesses" {
		response = &businesses.ResponseList{}
	} else if entity == "user" {
		response = &models.UserModel{}
	} else if entity == "project" {
		response = &models.ProjectModel{}
	} else if entity == "projects" {
		response = &projects.ResponseList{}
	} else if entity == "property" {
		response = &models.PropertyModel{}
	} else if entity == "properties" {
		response = &properties.ResponseList{}
	} else if entity == "person" {
		response = &models.PersonModel{}
	} else if entity == "people" {
		response = &people.ResponseList{}
	} else if entity == "invoice" {
		response = &models.InvoiceModel{}
	} else if entity == "invoices" {
		response = &invoices.ResponseList{}
	} else if entity == "blocks" {
		response = &blocks.ResponseList{}
	}

	err = json.Unmarshal(body, &response)
	return response, err
}

// GetResponseDelete success
func GetResponseDelete(res *http.Response) (fiber.Map, error) {
	response := fiber.Map{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}
