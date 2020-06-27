package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type AwsEnvConfig struct {
	AwsAccessKeyId     string `envconfig:"AWS_ACCESS_KEY_ID" required:"true"`
	AwsSecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY" required:"true"`
	AwsRegion          string `envconfig:"AWS_REGION" required:"true"`
	RdsUser            string `envconfig:"RDS_USER_NAME" default:"30" required:"true"`
	RdsPassword        string `envconfig:"RDS_PASSWORD" default:"30" required:"true"`
}

func ParseEnvConfiguration() (conf *AwsEnvConfig, err error) {
	conf = &AwsEnvConfig{}
	if err := envconfig.Process("", conf); err != nil {
		return nil, fmt.Errorf("Environment variables are not provided: %v ", err)
	}
	return
}
