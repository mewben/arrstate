package startup

import "go.mongodb.org/mongo-driver/mongo"

// Init -
func Init() *mongo.Database {
	InitEnvironment()
	db := ConnectMongo()
	// TODO:
	// Indexes(db)
	// Migrations(db)
	return db
}
