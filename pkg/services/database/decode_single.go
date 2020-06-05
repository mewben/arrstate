package database

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/models"
)

// DecodeSingle result
// Sets the cache on the found document
func DecodeSingle(result *mongo.SingleResult, collectionName string) interface{} {
	// TODO: for cache
	switch collectionName {
	case enums.CollBusinesses:
		business := models.NewBusinessModel()
		result.Decode(&business)
		return business

	case enums.CollUsers:
		user := models.NewUserModel()
		result.Decode(&user)
		return user

	case enums.CollPeople:
		person := models.NewPersonModel()
		result.Decode(&person)
		return person

	case enums.CollProjects:
		project := models.NewProjectModel()
		result.Decode(&project)
		return project

	case enums.CollLots:
		lot := models.NewLotModel()
		result.Decode(&lot)
		return lot
	}

	return nil

}
