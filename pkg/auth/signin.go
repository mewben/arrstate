package auth

import (
	"log"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// SigninPayload -
type SigninPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate Payload
func (v SigninPayload) Validate() error {
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

// Signin -
func (h *Handler) Signin(data *SigninPayload) (*models.AuthSuccessResponse, error) {
	log.Println("Signin")
	// validate payload
	if err := data.Validate(); err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	// get user by email
	filter := bson.D{
		{
			Key:   "email",
			Value: strings.ToLower(data.Email),
		},
	}
	userFound := h.DB.FindOne(h.Ctx, enums.CollUsers, filter)
	if userFound == nil {
		return nil, errors.NewHTTPError(errors.ErrSigninIncorrect)
	}

	// compare password
	user := userFound.(*models.UserModel)
	if !user.ComparePassword(user.Password, data.Password) {
		return nil, errors.NewHTTPError(errors.ErrSigninIncorrect)
	}

	// todo: get business domain by subdomain or whitelabel domain
	// for now, just get the the first id in person
	filter = bson.D{
		{
			Key:   "userID",
			Value: user.ID,
		},
	}
	personFound := h.DB.FindOne(h.Ctx, enums.CollPeople, filter)
	if personFound == nil {
		return nil, errors.NewHTTPError(errors.ErrSigninIncorrect)
	}
	person := personFound.(*models.PersonModel)

	response, err := h.AuthResponse(user.ID, person.BusinessID)
	if err != nil {
		log.Println("error authresponse", err)
		// TODO: remove business, user, and people
		return nil, err
	}
	// TODO: signin hooks

	return response, nil
}
