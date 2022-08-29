package cloud

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type mockSecretsManagerClient struct {
	secretsmanageriface.SecretsManagerAPI
}

// var getCountryFunc func(countryID string) (*locations.Country, *errors.APIerror)

func TestGetSecrets(t testing.T) {

}
