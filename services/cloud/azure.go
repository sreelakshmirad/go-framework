package cloud

type AzureConfig struct {
}

func (goaz *AzureConfig) GetSecrets(secretName interface{}) string {
	return "Secrets retreived succesfully"
}
