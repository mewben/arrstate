package people

import (
	"context"

	"github.com/mewben/realty278/pkg/services/database"
)

// Handler -
type Handler struct {
	DB  *database.Service
	Ctx context.Context
}
