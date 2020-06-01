package users

import (
	"errors"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/models"
)

// Payload -
type Payload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create User
func (h *Handler) Create(data *Payload) (*models.UserModel, error) {
	user := models.NewUserModel()
	hashedPassword, err := user.GeneratePassword(data.Password)
	if err != nil {
		log.Println("err generatepassword")
		return nil, err
	}
	user.Email = strings.ToLower(data.Email)
	user.Password = hashedPassword
	user.AccountStatus = enums.AccountStatusPending

	insertResult, err := h.DB.InsertOne(h.Ctx, enums.CollUsers, user)
	if err != nil {
		log.Println("err user.insertone")
		return nil, err
	}

	filter := bson.D{
		{
			Key:   "_id",
			Value: insertResult.InsertedID,
		},
	}
	insertedModel := h.DB.FindOne(h.Ctx, enums.CollUsers, filter)
	if insertedModel == nil {
		log.Println("err findone", insertedModel)
		return nil, errors.New("Not found")
	}

	return insertedModel.(*models.UserModel), nil

}
