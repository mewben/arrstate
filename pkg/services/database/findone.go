package database

import (
	"context"

	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOne - returns the document model
// caching is handled in the FindById
// but sets the cache after found
func (s *Service) FindOne(ctx context.Context, collectionName string, filter interface{}, opts ...*options.FindOneOptions) interface{} {
	result := s.DB.Collection(collectionName).FindOne(ctx, filter, opts...)

	if result.Err() != nil {
		return nil
	}

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
