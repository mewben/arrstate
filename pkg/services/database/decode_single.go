package database

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/errors"
	"github.com/mewben/realty278/pkg/models"
)

// DecodeSingle result
// Sets the cache on the found document
func DecodeSingle(result *mongo.SingleResult, collectionName string) (interface{}, error) {
	err := result.Err()
	// TODO: for cache
	switch collectionName {
	case enums.CollBusinesses:
		if err != nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundBusiness)
		}
		business := models.NewBusinessModel()
		result.Decode(&business)
		return business, nil

	case enums.CollUsers:
		if err != nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundUser)
		}
		user := models.NewUserModel()
		result.Decode(&user)
		return user, nil

	case enums.CollPeople:
		if err != nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundPerson)
		}
		person := models.NewPersonModel()
		result.Decode(&person)
		return person, nil

	case enums.CollProjects:
		if err != nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundProject)
		}
		project := models.NewProjectModel()
		result.Decode(&project)
		return project, nil

	case enums.CollProperties:
		if err != nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundProperty)
		}
		property := models.NewPropertyModel()
		result.Decode(&property)
		return property, nil

	case enums.CollInvoices:
		if err != nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundInvoice)
		}
		invoice := models.NewInvoiceModel()
		result.Decode(&invoice)
		return invoice, nil

	case enums.CollBlocks:
		if err != nil {
			return nil, errors.NewHTTPError(errors.ErrNotFoundBlock)
		}
		block := fiber.Map{}
		result.Decode(&block)
		return block, nil
	}

	return nil, errors.NewHTTPError(errors.ErrNotFound)

}
