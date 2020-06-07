package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/pkg/api/lots"
	"github.com/mewben/realty278/pkg/api/people"
	"github.com/mewben/realty278/pkg/api/projects"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// GetResponseError http
func GetResponseError(res *http.Response) (*errors.HTTPError, error) {
	responseError := &errors.HTTPError{}
	if res.StatusCode == 500 {
		msg, err := GetResponseMap(res)
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

// GetResponseMap general error
func GetResponseMap(res *http.Response) (string, error) {
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err

}

// GetResponseAuth success signup/signin
func GetResponseAuth(res *http.Response) (*models.AuthSuccessResponse, error) {
	response := &models.AuthSuccessResponse{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
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

// GetResponseProject success
func GetResponseProject(res *http.Response) (*models.ProjectModel, error) {
	response := &models.ProjectModel{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// GetResponseProjects success
func GetResponseProjects(res *http.Response) (*projects.ResponseList, error) {
	response := &projects.ResponseList{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// GetResponseLot success
func GetResponseLot(res *http.Response) (*models.LotModel, error) {
	response := &models.LotModel{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// GetResponseLots success
func GetResponseLots(res *http.Response) (*lots.ResponseList, error) {
	response := &lots.ResponseList{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// GetResponsePerson success
func GetResponsePerson(res *http.Response) (*models.PersonModel, error) {
	response := &models.PersonModel{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// GetResponsePersons success
func GetResponsePersons(res *http.Response) (*people.ResponseList, error) {
	response := &people.ResponseList{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}
