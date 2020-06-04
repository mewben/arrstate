package auth

import (
	"log"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/api/businesses"
	"github.com/mewben/realty278/pkg/api/people"
	"github.com/mewben/realty278/pkg/api/users"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// SignupPayload -
type SignupPayload struct {
	GivenName  string `json:"givenName" validate:"required"`
	FamilyName string `json:"familyName"`
	Business   string `json:"business" validate:"required"`
	Domain     string `json:"domain" validate:"required"`
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required,min=6"`
}

// Signup -
func (h *Handler) Signup(data *SignupPayload) (*models.AuthSuccessResponse, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
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

	response, err := h.AuthResponse(user.ID, business.ID)
	if err != nil {
		log.Println("error authresponse", err)
		// TODO: remove business, user, and people
		return nil, err
	}

	return response, nil
}
