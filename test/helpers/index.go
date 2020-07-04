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
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// CheckJWT and returns userID and businessID parsed from the token
func CheckJWT(token string, assert *assert.Assertions) (userID, businessID primitive.ObjectID) {
	assert.NotEmpty(token)
	tokenSigningKey := viper.GetString("TOKEN_SIGNING_KEY")
	assert.NotEmpty(tokenSigningKey)

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenSigningKey), nil
	})
	assert.Nil(err, t)

	// assert claims
	claims := t.Claims.(jwt.MapClaims)
	exp := time.Now().Add(time.Hour * viper.GetDuration("TOKEN_EXPIRY")).Unix()
	claimsExpiry := claims["exp"].(float64)
	diff := float64(exp) - claimsExpiry
	assert.LessOrEqual(diff, float64(1))
	// userID
	userID, err = primitive.ObjectIDFromHex(claims["sub"].(string))
	assert.Nil(err)
	assert.False(userID.IsZero())

	// businessID
	businessID, err = primitive.ObjectIDFromHex(claims["businessID"].(string))
	assert.Nil(err)
	assert.False(businessID.IsZero())
	return
}
