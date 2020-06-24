package models

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// UserModel -
type UserModel struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email         string             `bson:"email" json:"email"`
	Password      string             `bson:"password" json:"-"`
	AccountStatus string             `bson:"accountStatus" json:"accountStatus"`
}

// CurrentUser -
type CurrentUser struct {
	User   *UserModel   `json:"user"`
	Person *PersonModel `json:"person"`
}

// MeModel -
type MeModel struct {
	CurrentUser     *CurrentUser     `json:"currentUser"`
	CurrentBusiness *BusinessModel   `json:"currentBusiness"`
	UserBusinesses  []*BusinessModel `json:"userBusinesses"`
}

// NewUserModel -
func NewUserModel() *UserModel {
	return &UserModel{}
}

// NewCurrentUser -
func NewCurrentUser(userID, businessID primitive.ObjectID) *CurrentUser {
	return &CurrentUser{
		User:   NewUserModel(),
		Person: NewPersonModel(userID, businessID),
	}
}

// NewMeModel -
func NewMeModel(userID, businessID primitive.ObjectID) *MeModel {
	return &MeModel{
		CurrentUser:     NewCurrentUser(userID, businessID),
		CurrentBusiness: NewBusinessModel(),
		UserBusinesses:  make([]*BusinessModel, 0),
	}
}

// GenerateJWT for signup/signin repsonse
func (*UserModel) GenerateJWT(sub, businessID primitive.ObjectID) (string, error) {
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
