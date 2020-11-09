package startup

import (
	"github.com/markbates/pkger"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mewben/arrstate/internal/enums"
)

// Init -
func Init() *mongo.Database {
	// pkger include assets
	pkger.Include("/assets/countries.json")
	pkger.Include("/assets/currencies.json")
	pkger.Include("/web/static/locales/en/global.json")

	InitEnvironment()
	db := ConnectMongo()
	Indexes(db)
	enums.Init()
	// TODO:
	// Migrations(db)

	// cron.Run()

	return db
}
