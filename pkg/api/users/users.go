package users

import (
	"context"

	validator "github.com/go-playground/validator/v10"

	"github.com/mewben/realty278/pkg/services/database"
)

// Handler -
type Handler struct {
	DB  *database.Service
	Ctx context.Context
}

// Payload -
type Payload struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=6"`
}

// use a single instance of Validate, it caches struct info
var validate = validator.New()
