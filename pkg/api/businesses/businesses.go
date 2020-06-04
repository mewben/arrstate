package businesses

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

// Params -
type Params struct{}

// use a single instance of Validate, it caches struct info
var validate = validator.New()
