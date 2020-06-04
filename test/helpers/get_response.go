package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// GetResponseError http
func GetResponseError(res *http.Response) (*errors.HTTPError, error) {
	responseError := &errors.HTTPError{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return responseError, nil
	}
	err = json.Unmarshal(body, &responseError)
	return responseError, err
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
