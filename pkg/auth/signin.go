package auth

import (
	"log"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// SigninPayload -
type SigninPayload struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`
}

// Signin -
func (h *Handler) Signin(data *SigninPayload) (fiber.Map, error) {
	log.Println("Signin")
	// validate payload
	if err := validate.Struct(data); err != nil {
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

	token, err := user.GenerateJWT(user.ID, person.BusinessID)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrSigninIncorrect, err)
	}

	// TODO: signin hooks

	return fiber.Map{"token": token}, nil
}
