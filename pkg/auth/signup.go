package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/gommon/random"
	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/api/businesses"
	"github.com/mewben/arrstate/pkg/api/people"
	"github.com/mewben/arrstate/pkg/api/users"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
)

// SignupPayload -
type SignupPayload struct {
	Name     models.PersonName `json:"name"`
	Business string            `json:"business" validate:"required"`
	Domain   string            `json:"domain" validate:"required"`
	Email    string            `json:"email" validate:"email,required"`
	Password string            `json:"password" validate:"required,min=6"`
}

// Signup -
func (h *Handler) Signup(data *SignupPayload) (fiber.Map, error) {
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
		BusinessModel: models.BusinessModel{
			Name:   data.Business,
			Domain: data.Domain,
		},
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
		Email:      data.Email,
		Password:   data.Password,
		DeviceCode: random.String(32),
	}
	user, err := userHandler.Create(userPayload)
	if err != nil {
		log.Println("err create user", err)
		// TODO: remove business
		return nil, err
	}

	// 3. Create Person
	personHandler := &people.Handler{
		DB:       h.DB,
		Ctx:      h.Ctx,
		User:     user,
		Business: business,
	}
	personPayload := &people.Payload{
		PersonModel: models.PersonModel{
			UserID: &user.ID,
			Email:  data.Email,
			Role:   []string{enums.RoleOwner},
			Name:   data.Name,
		},
	}
	_, err = personHandler.Create(personPayload)
	if err != nil {
		log.Println("error create person", err)
		// TODO: remove business and user
		return nil, err
	}

	// TODO: signup hooks
	return fiber.Map{"deviceCode": user.DeviceCode, "domain": business.Domain}, nil
}
