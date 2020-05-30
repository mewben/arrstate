package internal

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

// InitEnvironment reads environment variables
// to be accessed throughout the whole app
func InitEnvironment() {
	env := os.Getenv("ENV")

	if env == "PRODUCTION" {
		viper.AutomaticEnv()
		return
	}

	configFile := ".env.development"
	if env == "TESTING" {
		configFile = ".env.testing"
	}
	// set the configPath to root directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No caller information")
	}
	viper.AddConfigPath(path.Join(path.Dir(filename), "../"))
	viper.SetConfigType("env")
	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Environment error", err)
	}
}
