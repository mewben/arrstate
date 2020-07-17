package blocks

import (
	"context"

	"github.com/mewben/realty278/pkg/models"
	"github.com/mewben/realty278/pkg/services/database"
)

// Handler -
type Handler struct {
	DB       *database.Service
	Ctx      context.Context
	User     *models.UserModel
	Business *models.BusinessModel
}
