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
func (h *Handler) Signin(data *SigninPayload, domain string) (fiber.Map, error) {
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

	// get business by domain
	// TODO: whitelabel domain
	filter = bson.D{
		{
			Key:   "domain",
			Value: domain,
		},
	}
	businessFound := h.DB.FindOne(h.Ctx, enums.CollBusinesses, filter)
	if businessFound == nil {
		return nil, errors.NewHTTPError(errors.ErrNotFoundBusiness)
	}
	business := businessFound.(*models.BusinessModel)

	filter = bson.D{
		{
			Key:   "userID",
			Value: user.ID,
		},
		{
			Key:   "businessID",
			Value: business.ID,
		},
	}
	personFound := h.DB.FindOne(h.Ctx, enums.CollPeople, filter)
	if personFound == nil {
		return nil, errors.NewHTTPError(errors.ErrUserNotInBusiness)
	}
	person := personFound.(*models.PersonModel)

	token, err := user.GenerateJWT(user.ID, person.BusinessID)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrSigninIncorrect, err)
	}

	// TODO: signin hooks

	return fiber.Map{"token": token}, nil
}
