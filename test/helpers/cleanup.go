package helpers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/internal/enums"
)

// CleanupFixture -
func CleanupFixture(db *mongo.Database) {
	collections := []string{
		enums.CollBusinesses,
		enums.CollUsers,
		enums.CollPeople,
		enums.CollProjects,
		enums.CollProperties,
		enums.CollInvoices,
		enums.CollBlocks,
		enums.CollFiles,
	}
	for _, col := range collections {
		db.Collection(col).DeleteMany(context.TODO(), bson.D{})
	}
}
