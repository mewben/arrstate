package startup

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo -
// MONGO_URL=mongodb://root:password@localhost:27017/databaseName?authSource=admin
func ConnectMongo() *mongo.Database {
	mongoURL := viper.GetString("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalln("Error mongo NewClient:", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln("Error mongo Connect:", err)
	}

	return client.Database(extractDatabaseName(mongoURL))

}

// extractDatabaseName from the uri
func extractDatabaseName(mongoURL string) (databaseName string) {
	databaseName = "default"

	splits := strings.Split(mongoURL, "/")
	if len(splits) < 4 {
		return
	}

	splits = strings.Split(splits[3], "?")
	if splits[0] == "" {
		return
	}

	return splits[0]
}
