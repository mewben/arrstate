package auth

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// SigninPayload -
type SigninPayload struct {
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required,min=6"`
	GrantType  string `json:"grant_type"`
	DeviceCode string `json:"deviceCode"`
}

// Signin -
func (h *Handler) Signin(data *SigninPayload, domain string) (*SigninResponse, error) {
	user := &models.UserModel{}
	if data.GrantType == "device_code" && data.DeviceCode != "" {
		// find user by device_code
		filter := bson.M{"deviceCode": data.DeviceCode}
		userFound, _ := h.DB.FindOne(h.Ctx, enums.CollUsers, filter)
		if userFound == nil {
			return nil, errors.NewHTTPError(errors.ErrSigninIncorrect)
		}
		user = userFound.(*models.UserModel)

	} else {
		// validate payload
		if err := validate.Struct(data); err != nil {
			return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
		}

		// get user by email
		filter := bson.M{"email": strings.ToLower(data.Email)}
		userFound, _ := h.DB.FindOne(h.Ctx, enums.CollUsers, filter)
		if userFound == nil {
			return nil, errors.NewHTTPError(errors.ErrSigninIncorrect)
		}
		user = userFound.(*models.UserModel)
		// compare password
		if !user.ComparePassword(user.Password, data.Password) {
			return nil, errors.NewHTTPError(errors.ErrSigninIncorrect)
		}

	}

	// get business by domain
	// TODO: whitelabel domain
	filter := bson.D{
		{
			Key:   "domain",
			Value: domain,
		},
	}
	businessFound, err := h.DB.FindOne(h.Ctx, enums.CollBusinesses, filter)
	if err != nil {
		return nil, err
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
	personFound, _ := h.DB.FindOne(h.Ctx, enums.CollPeople, filter)
	if personFound == nil {
		return nil, errors.NewHTTPError(errors.ErrUserNotInBusiness)
	}
	person := personFound.(*models.PersonModel)

	token, err := user.GenerateJWT(user.ID, person.BusinessID)
	if err != nil {
		return nil, errors.NewHTTPError(errors.ErrSigninIncorrect, err)
	}

	go h.SigninHook(user)

	response := &SigninResponse{
		Token: token,
		User:  user,
	}

	return response, nil
}
