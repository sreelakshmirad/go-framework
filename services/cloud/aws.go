package cloud

import (
	"encoding/base64"
	"fmt"
	"framework_v1/core/config"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/mitchellh/mapstructure"
)

type AWSConfig struct {
	AwsEndpoint   string
	AwsRegion     string
	AwsSecretname string
	AwsCurrent    string
	AwsProfile    string
}

func (awsConfigs *AWSConfig) GetSecrets(secretName interface{}) string {
	var decodedSecretString string

	// fmt.Println(goaws.AwsSecretname, "hello")

	// var awsConfigs AWSConfig
	err := mapstructure.Decode(config.CloudConfig, &awsConfigs)
	if err != nil {
		log.Fatal(err)
	}

	secretNameStr, ok := secretName.(string)
	if ok {
		awsConfigs.AwsSecretname = secretNameStr
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsConfigs.AwsRegion),
		Credentials: credentials.NewSharedCredentials("", awsConfigs.AwsProfile),
		Endpoint:    aws.String(awsConfigs.AwsEndpoint),
	},
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	}

	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{

		SecretId: aws.String(awsConfigs.AwsSecretname),
	}
	fmt.Println(input, "input")
	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.Fatal(err)

	} else {
		// var secretString string
		// fmt.Println(result)

		if result.SecretString != nil {
			decodedSecretString = *result.SecretString
			// json.Unmarshal([]byte(secretString), &decodedSecretString)
			// fmt.Println(secretString, "secretString")
		} else {
			decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
			len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
			if err != nil {
				fmt.Println("Base64 Decode Error:", err)
			}
			decodedSecretString = string(decodedBinarySecretBytes[:len])
			// json.Unmarshal([]byte(decodedBinarySecret), &decodedSecretString)
		}
	}

	// fmt.Println(decodedSecretString, "decodedSecretString")
	return decodedSecretString

}
