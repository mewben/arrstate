package users

import (
	"log"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Payload -
type Payload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate Payload
func (v Payload) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(
			&v.Email,
			validation.Required,
			is.EmailFormat,
		),
		validation.Field(
			&v.Password,
			validation.Required,
			validation.Length(6, 0),
		),
	)
}

// Create User
func (h *Handler) Create(data *Payload) (*models.UserModel, error) {
	// validate payload
	if err := data.Validate(); err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	cleanedEmail := strings.ToLower(data.Email)

	// check duplicate user email
	filter := bson.D{
		{
			Key:   "email",
			Value: cleanedEmail,
		},
	}
	userFound := h.DB.FindOne(h.Ctx, enums.CollUsers, filter)
	if userFound != nil {
		return nil, errors.NewHTTPError(errors.ErrUserDuplicate)
	}

	user := models.NewUserModel()
	hashedPassword, err := user.GeneratePassword(data.Password)
	if err != nil {
		log.Println("err generatepassword")
		return nil, err
	}
	user.Email = cleanedEmail
	user.Password = hashedPassword
	user.AccountStatus = enums.AccountStatusPending

	insertResult, err := h.DB.InsertOne(h.Ctx, enums.CollUsers, user)
	if err != nil {
		log.Println("err user.insertone")
		return nil, err
	}

	filter = bson.D{
		{
			Key:   "_id",
			Value: insertResult.InsertedID,
		},
	}
	insertedModel := h.DB.FindOne(h.Ctx, enums.CollUsers, filter)
	if insertedModel == nil {
		log.Println("err findone", insertedModel)
		return nil, errors.NewHTTPError(errors.ErrNotFound)
	}

	return insertedModel.(*models.UserModel), nil

}
