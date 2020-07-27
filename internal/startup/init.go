package startup

import (
	"github.com/mewben/arrstate/internal/enums"
	"go.mongodb.org/mongo-driver/mongo"
)

// Init -
func Init() *mongo.Database {
	InitEnvironment()
	db := ConnectMongo()
	Indexes(db)
	enums.Init()
	// TODO:
	// Migrations(db)
	return db
}
