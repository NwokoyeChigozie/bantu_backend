package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

type ServerConfiguration struct {
	Port                      string
	Secret                    string
	AccessTokenExpireDuration int
}

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// Setup initialize configuration
var (
	Config   *Configuration
)


func Setup() {
	var configuration *Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	fmt.Println(viper.Get("DB_PORT"))

	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file, %s", err)
	// }

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// Params = configuration.Params
	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
