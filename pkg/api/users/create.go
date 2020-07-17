package users

import (
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// Create User
func (h *Handler) Create(data *Payload) (*models.UserModel, error) {
	// validate payload
	if err := validate.Struct(data); err != nil {
		return nil, errors.NewHTTPError(errors.ErrInputInvalid, err)
	}

	cleanedEmail := strings.ToLower(data.Email)

	// check duplicate user email
	filter := bson.D{
		{
			Key:   "email",
			Value: cleanedEmail,
		},
	}
	userFound, _ := h.DB.FindOne(h.Ctx, enums.CollUsers, filter)
	if userFound != nil {
		return nil, errors.NewHTTPError(errors.ErrUserDuplicate)
	}

	user := models.NewUserModel()
	hashedPassword, err := user.GeneratePassword(data.Password)
	if err != nil {
		log.Println("err generatepassword")
		return nil, err
	}
	user.Email = cleanedEmail
	user.Password = hashedPassword
	user.AccountStatus = enums.AccountStatusActive
	user.DeviceCode = data.DeviceCode

	doc, err := h.DB.InsertOne(h.Ctx, enums.CollUsers, user)
	if err != nil || doc == nil {
		log.Println("insertonerr", err)
		return nil, errors.NewHTTPError(errors.ErrInsert, err)
	}

	user = doc.(*models.UserModel)

	return user, nil

}
