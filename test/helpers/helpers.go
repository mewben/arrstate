package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// DoRequest helper
func DoRequest(method, path string, body interface{}) *http.Request {
	var payload io.Reader
	bodyString := ""
	if body != nil {
		// encode body to json
		jsonByte, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		bodyString = string(jsonByte)
		payload = strings.NewReader(bodyString)
	}
	req, err := http.NewRequest(method, path, payload)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)
	req.Header.Set("Content-Length", strconv.Itoa(len(bodyString)))
	return req
}

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
		return response, nil
	}
	err = json.Unmarshal(body, &response)
	return response, err
}
