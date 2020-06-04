package models

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// UserModel -
type UserModel struct {
	ID            string `bson:"_id,omitempty" json:"_id,omitempty"`
	Email         string `bson:"email" json:"email"`
	Password      string `bson:"password" json:"-"`
	AccountStatus string `bson:"accountStatus" json:"accountStatus"`
}

// NewUserModel -
func NewUserModel() *UserModel {
	return &UserModel{}
}

// GenerateJWT for signup/signin repsonse
func (*UserModel) GenerateJWT(sub, businessID string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = sub
	claims["businessID"] = businessID
	claims["exp"] = time.Now().Add(time.Hour * viper.GetDuration("TOKEN_EXPIRY")).Unix()

	// Generate encoded token
	return token.SignedString([]byte(viper.GetString("TOKEN_SIGNING_KEY")))
}

// GeneratePassword -
func (*UserModel) GeneratePassword(pwd string) (string, error) {
	pwd = pwd + viper.GetString("PASSWORD_SALT")
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePassword -
func (*UserModel) ComparePassword(hashedPassword, plainPassword string) bool {
	plainPassword = plainPassword + viper.GetString("PASSWORD_SALT")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		log.Println("error ComparePassword", err)
		return false
	}

	return true
}
