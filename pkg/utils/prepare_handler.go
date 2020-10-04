package utils

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mewben/arrstate/internal/enums"
	"github.com/mewben/arrstate/pkg/errors"
	"github.com/mewben/arrstate/pkg/models"
	"github.com/mewben/arrstate/pkg/services/database"
)

// PrepareHandler -
func PrepareHandler(c *fiber.Ctx, db *database.Service) (user *models.UserModel, business *models.BusinessModel, person *models.PersonModel, err error) {
	userID, businessID, personID, err := ExtractClaims(c)
	if err != nil {
		return
	}

	// get user
	user = db.FindUser(c.Context(), userID)
	if user == nil {
		err = errors.NewHTTPError(errors.ErrNotFoundUser)
		return
	}

	// get business
	business = db.FindBusiness(c.Context(), businessID)
	if business == nil {
		err = errors.NewHTTPError(errors.ErrNotFoundBusiness)
		return
	}

	// get person
	personFound, _ := db.FindByID(c.Context(), enums.CollPeople, personID, businessID)
	if personFound == nil {
		err = errors.NewHTTPError(errors.ErrNotFoundPerson)
		return
	}
	person = personFound.(*models.PersonModel)

	return
}
