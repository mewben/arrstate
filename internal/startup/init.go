package startup

import "go.mongodb.org/mongo-driver/mongo"

// Init -
func Init() *mongo.Database {
	InitEnvironment()
	db := ConnectMongo()
	Indexes(db)
	// TODO:
	// Migrations(db)
	return db
}
