package cloud

type CloudConnection interface {
	GetSecrets(interface{}) string
}

func NewCloudConnector(driver string) (castType CloudConnection) {
	castType = &AWSConfig{}
	switch driver {
	case "aws":
		castType = &AWSConfig{}
	case "azure":
		castType = &AzureConfig{}
	default:
		castType = &AWSConfig{}

	}
	return
}
