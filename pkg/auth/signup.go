package auth

import (
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/api/businesses"
	"github.com/mewben/realty278/pkg/api/people"
	"github.com/mewben/realty278/pkg/api/users"
	"github.com/mewben/realty278/pkg/errors"
)

// SignupPayload -
type SignupPayload struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Business   string `json:"business"`
	Domain     string `json:"domain"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

// Validate Payload
func (v SignupPayload) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(
			&v.GivenName,
			validation.Required,
		),
		validation.Field(
			&v.Business,
			validation.Required,
		),
		validation.Field(
			&v.Domain,
			validation.Required,
		),
		validation.Field(
			&v.Email,
			validation.Required,
		),
		validation.Field(
			&v.Password,
			validation.Required,
		),
	)
}

// Signup -
func (h *Handler) Signup(data *SignupPayload) (interface{}, error) {
	log.Println("Signup")
	// validate payload
	if err := data.Validate(); err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid)
	}

	// 1. Create Business
	businessHandler := &businesses.Handler{
		DB:  h.DB,
		Ctx: h.Ctx,
	}
	businessPayload := &businesses.Payload{
		Name:   data.Business,
		Domain: data.Domain,
	}
	business, err := businessHandler.Create(businessPayload)
	if err != nil {
		log.Println("create business", business)
		return nil, err
	}

	log.Println("after create business")

	// 2. Create User
	userHandler := &users.Handler{
		DB:  h.DB,
		Ctx: h.Ctx,
	}
	userPayload := &users.Payload{
		Email:    data.Email,
		Password: data.Password,
	}
	user, err := userHandler.Create(userPayload)
	if err != nil {
		log.Println("err create user", err)
		// TODO: remove business
		return nil, err
	}

	log.Println("after create user")

	// 3. Create Person
	personHandler := &people.Handler{
		DB:  h.DB,
		Ctx: h.Ctx,
	}
	personPayload := &people.Payload{
		UserID:     user.ID,
		BusinessID: business.ID,
		Role:       enums.RoleOwner,
		GivenName:  data.GivenName,
		FamilyName: data.FamilyName,
	}
	_, err = personHandler.Create(personPayload)
	if err != nil {
		log.Println("error create person", err)
		// TODO: remove business and user
		return nil, err
	}

	log.Println("afeter create person")

	response, err := h.AuthResponse(user.ID, business.ID)
	if err != nil {
		log.Println("error authresponse", err)
		// TODO: remove business, user, and people
		return nil, err
	}

	log.Println("after authresponse")

	return response, nil
}
