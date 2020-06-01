package startup

import (
	"context"
	"log"

	"github.com/mewben/realty278/internal/enums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Indexes -
func Indexes(db *mongo.Database) {
	var err error
	// Businesses
	_, err = db.Collection(enums.CollBusinesses).Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   "domain",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Fatalln("error business index", err)
	}

	// Users
	_, err = db.Collection(enums.CollUsers).Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   "email",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Fatalln("error user index", err)
	}
}
