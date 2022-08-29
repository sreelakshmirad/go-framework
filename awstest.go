package main

import (
	// "aws-secret-manager-test/Logger"
	// "aws-secret-manager-test/Models"
	"database/sql"
	"fmt"
	"framework_v1/internal/helpers"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type ClientDatabase struct {
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string
}

func main() {

	// fmt.Println(config.CloudConfig)()
	helpers.GetSecretsRetriever()
	os.Exit(8)

	// helpers.GetCreateDBConnection()
	// getDatabaseAuth()

	fmt.Println("Hello Go")
	// psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	databaseAuth.Host, databaseAuth.Port, databaseAuth.UserName, databaseAuth.Password, databaseAuth.DBName)

	// DB, err = sql.Open("postgres", psql)
	// if err != nil {
	// 	fmt.Println(err.Error())

	// }
	// fmt.Println(DB)
	// fmt.Println()
	// if err != nil {
	// 	Logger.AddLogger(Logger.ERROR, "Database driver error")
	// 	panic(err)

	// }
	// if err = DB.Ping(); err != nil {
	// 	Logger.AddLogger(Logger.ERROR, "Database parameters error")
	// 	panic(err)
	// }
	// Logger.AddLogger(Logger.INFO, "Connected to Database")
}

// func getDatabaseAuth() godb.DBConnector {

// 	// var databaseAuth string
// 	var secretInterface interface{}

// 	sess, err := session.NewSession(&aws.Config{
// 		Region:      aws.String("us-east-1"),
// 		Credentials: credentials.NewSharedCredentials("", "localstack"),
// 		Endpoint:    aws.String("http://localhost:4566"),
// 	},
// 	)

// 	_, err = sess.Config.Credentials.Get()

// 	svc := secretsmanager.New(sess)
// 	input := &secretsmanager.GetSecretValueInput{

// 		SecretId: aws.String("client1/postgres-cred"),
// 	}

// 	result, err := svc.GetSecretValue(input)
// 	fmt.Println(err)
// 	fmt.Println(result, "result")

// 	if err == nil {
// 		var secretString, decodedBinarySecret string

// 		fmt.Println(result)

// 		if result.SecretString != nil {
// 			secretString = *result.SecretString

// 			fmt.Println(secretString, "secretString")
// 			json.Unmarshal([]byte(secretString), &secretInterface)
// 			fmt.Println(secretInterface, "secretInterface")
// 			json.Unmarshal([]byte(secretString), &secretInterface)

// 			// change here

// 			// var dbStruct cloud.AWSConfig

// 			dbStruct := godb.NewDbConnector("postgres")

// 			err := mapstructure.Decode(secretInterface, &dbStruct)
// 			fmt.Println(err)
// 			db, err := dbStruct.CreateDBConnection()
// 			fmt.Println(db)

// 			// err := json.Unmarshal([]byte(secretString), &dbStruct)
// 			fmt.Println(dbStruct, "dbStruct")
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		} else {
// 			decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
// 			len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
// 			if err != nil {
// 				fmt.Println("Base64 Decode Error:", err)
// 			}
// 			decodedBinarySecret = string(decodedBinarySecretBytes[:len])
// 			json.Unmarshal([]byte(decodedBinarySecret), &secretInterface)

// 			json.Unmarshal([]byte(secretString), &secretInterface)

// 		}
// 	}

// 	// change here

// 	// var dbStruct cloud.AWSConfig

// }
