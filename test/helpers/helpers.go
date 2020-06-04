package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/pkg/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// DoRequest helper
func DoRequest(method, path string, body interface{}, token string) *http.Request {
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
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return req
}

// CheckJWT -
func CheckJWT(token string, user *models.UserModel, businessID string, assert *assert.Assertions) {
	assert.NotEmpty(token)
	tokenSigningKey := viper.GetString("TOKEN_SIGNING_KEY")
	assert.NotEmpty(tokenSigningKey)
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenSigningKey), nil
	})
	assert.Nil(err, t)
	claims := t.Claims.(jwt.MapClaims)
	exp := time.Now().Add(time.Hour * viper.GetDuration("TOKEN_EXPIRY")).Unix()
	claimsExpiry := claims["exp"].(float64)
	diff := float64(exp) - claimsExpiry
	assert.Equal(user.ID, claims["sub"])
	assert.LessOrEqual(diff, float64(1))
	assert.Equal(businessID, claims["businessID"])
}
