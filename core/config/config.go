package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	// General Configurations
	AppPortEvar = "APP_PORT"
	LogFileEvar = "LOGGER_FILE"

	// Cloud Configuration
	CloudProviderEvar = "CLOUD_PROVIDER"

	// SQL ENV Variables
	// #TODO
	DBUserEvar   = "DB_USER"   // username
	DBPassEvar   = "DB_PASS"   // password
	DBHostEvar   = "DB_HOST"   // host
	DBPortEvar   = "DB_PORT"   // port
	DBDbEvar     = "DB_NAME"   // db name
	DBDriverEvar = "DB_DRIVER" // db name

)

// Exported  Global Variables
var (
	ClientConfigMap map[string]interface{}
	CloudConfig     interface{}
)

// Configurations
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
	Logger   LoggingConfigurations
	Cloud    CloudConfigurations
}

// ServerConfigurations
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations
type DatabaseConfigurations struct {
	DBDriver   string
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
}

type LoggingConfigurations struct {
	Loggerfile string
}
type CloudValues struct {
	CloudVal map[string]map[string]string
}
type CloudConfigurations struct {
	Provider      string
	CloudSettings map[string]interface{} `mapstructure:"cloud_settings"`
}

func init() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	var config Configurations

	err := viper.Unmarshal(&config)

	if err != nil {
		log.Fatal(err)
	}

	CloudConfig = config.Cloud.CloudSettings[config.Cloud.Provider]

	// fmt.Println(CloudConfig, "CloudConfigCloudConfigCloudConfig")
	viper.SetConfigName("db-config")
	viper.AddConfigPath(".")

	viper.SetConfigType("yml")
	viper.MergeInConfig()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&ClientConfigMap)

	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	// Server Configurations
	viper.SetDefault(AppPortEvar, config.Server.Port)
	viper.BindEnv(AppPortEvar)

	// Logging Configurations
	viper.SetDefault(LogFileEvar, config.Logger.Loggerfile)
	viper.BindEnv(LogFileEvar)

	// Cloud Configurations

	viper.SetDefault(CloudProviderEvar, config.Cloud.Provider)
	viper.BindEnv(CloudProviderEvar)

	viper.SetDefault(DBUserEvar, config.Database.DBUser)
	viper.BindEnv(DBUserEvar)

	viper.SetDefault(DBPassEvar, config.Database.DBPassword)
	viper.BindEnv(DBPassEvar)

	viper.SetDefault(DBHostEvar, config.Database.DBHost)
	viper.BindEnv(DBHostEvar)

	viper.SetDefault(DBPortEvar, config.Database.DBPort)
	viper.BindEnv(DBPortEvar)

	viper.SetDefault(DBDbEvar, config.Database.DBName)
	viper.BindEnv(DBDbEvar)

	viper.SetDefault(DBDriverEvar, config.Database.DBDriver)
	viper.BindEnv(DBDriverEvar)

}
