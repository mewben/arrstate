package auth

import (
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/models"
)

// SigninHook -
func (h *Handler) SigninHook(user *models.UserModel) {
	if user.DeviceCode != "" {
		// remove deviceCode
		filter := bson.M{"_id": user.ID}
		op := bson.D{
			{
				Key:   "$unset",
				Value: bson.M{"deviceCode": ""},
			},
		}

		h.DB.FindOneAndUpdate(h.Ctx, enums.CollUsers, filter, op)
	}
}
