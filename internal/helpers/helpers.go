package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"framework_v1/core/config"
	"framework_v1/services/cloud"
	"framework_v1/services/godb"

	"log"
	"regexp"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// helper function to prepare a mysql db connection
func MustReadMysqlConf() string {
	// driver := viper.GetString(config.DBDriverEvar)

	usr := viper.GetString(config.DBUserEvar)
	pwd := viper.GetString(config.DBPassEvar)
	host := viper.GetString(config.DBHostEvar)
	port := viper.GetString(config.DBPortEvar)
	db := viper.GetString(config.DBDbEvar)

	if len(usr) == 0 {
		log.Fatal("invalid MYSQL_USER")
	}
	if len(pwd) == 0 {
		log.Fatal("invalid MYSQL_PASSWORD")
	}
	CreateDBConnection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		usr,
		pwd,
		host,
		port,
		db,
	)

	return CreateDBConnection
}

// regular expression
var scrubRex = regexp.MustCompile(`:[^:]+@`)

// MustPrepareDB
// Wrapper method for db connection
// @param : connString string; mysql connection string
// @return: *sql.DB connection
func MustPrepareDB(connString string) *sql.DB {
	db, err := ConnectMySQL(connString)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// ConnectMySQL
// Attempts a db connection with prepared connection string
// @param : connString string; mysql connection string
// @return: *sql.DB connection, error
func ConnectMySQL(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrapf(err,
			"failed to connect to %s",
			ScrubMySQLPassword(connString),
		)
	}
	return db, nil
}

// ScrubMySQLPassword
// hide password from the output streams
// @param: cs string connection string
// @return: string
func ScrubMySQLPassword(cs string) string {
	return scrubRex.ReplaceAllString(cs, ":PWD@")
}

func GetCreateDBConnection() {
	dbConfigs := config.ClientConfigMap

	for _, configs := range dbConfigs {
		result := godb.NewDbConnector(viper.GetString("DB_DRIVER"))
		err := mapstructure.Decode(configs, &result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result, "resultresult")
		db, err := result.CreateDBConnection()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(db)

	}

}

func GetSecretsRetriever() {
	clientName := "folium"
	cloudConfigs := config.CloudConfig
	dbConfigs := config.ClientConfigMap["db_settings"]

	result := cloud.NewCloudConnector(viper.GetString("CLOUD_PROVIDER"))

	err := mapstructure.Decode(cloudConfigs, &result)
	if err != nil {
		log.Fatal(err)
	}

	var dbConfigsMap map[string]string
	mapstructure.Decode(dbConfigs, &dbConfigsMap)
	secretString := fmt.Sprintf(dbConfigsMap["secret"], clientName)

	clientdbConfigs := result.GetSecrets(secretString)

	fmt.Println(clientdbConfigs, "clientdbConfigs")
	var secretInterface interface{}

	json.Unmarshal([]byte(clientdbConfigs), &secretInterface)

	dbStruct := godb.NewDbConnector("postgres")

	err = mapstructure.Decode(secretInterface, &dbStruct)
	if err != nil {
		log.Fatal(err)
	}
	db, err := dbStruct.CreateDBConnection()
	fmt.Println(db)

	if err != nil {
		log.Fatal(err)
	}

}
