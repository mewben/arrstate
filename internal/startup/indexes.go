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
		},
		{
			Keys: bson.D{
				{
					Key:   "deviceCode",
					Value: 1,
				},
			},
		},
	})
	if err != nil {
		log.Fatalln("error user index", err)
	}

	// People
	_, err = db.Collection(enums.CollPeople).Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   "userID",
					Value: 1,
				},
			},
		},
		{
			Keys: bson.D{
				{
					Key:   "businessID",
					Value: 1,
				},
			},
		},
		{
			Keys: bson.D{
				{
					Key:   "email",
					Value: 1,
				},
				{
					Key:   "businessID",
					Value: 1,
				},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{
					Key:   "role",
					Value: 1,
				},
				{
					Key:   "businessID",
					Value: 1,
				},
			},
		},
	})
	if err != nil {
		log.Fatalln("error people index", err)
	}

	// Projects
	_, err = db.Collection(enums.CollProjects).Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   "userID",
					Value: 1,
				},
			},
		},
		{
			Keys: bson.D{
				{
					Key:   "businessID",
					Value: 1,
				},
			},
		},
	})
	if err != nil {
		log.Fatalln("error project index", err)
	}

	// Properties
	_, err = db.Collection(enums.CollProperties).Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   "projectID",
					Value: 1,
				},
			},
		},
		{
			Keys: bson.D{
				{
					Key:   "businessID",
					Value: 1,
				},
			},
		},
	})
	if err != nil {
		log.Fatalln("error project index", err)
	}

	// ClientProperties
	_, err = db.Collection(enums.CollClientProperties).Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{
					Key:   "propertyID",
					Value: 1,
				},
			},
		},
		{
			Keys: bson.D{
				{
					Key:   "businessID",
					Value: 1,
				},
			},
		},
	})
	if err != nil {
		log.Fatalln("error clientproperties index", err)
	}

}
