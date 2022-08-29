package logger

import (
	"os"

	"framework_v1/core/config"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	loggerFile := viper.GetString(config.LogFileEvar)

	f, err := os.OpenFile(loggerFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// Output to a file 'error.log'
	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
