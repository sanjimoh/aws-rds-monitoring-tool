package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type AwsEnvConfig struct {
	AwsAccessKeyId     string `envconfig:"AWS_ACCESS_KEY_ID" required:"true"`
	AwsSecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY" required:"true"`
	AwsSessionToken    string `envconfig:"AWS_SESSION_TOKEN" default:"30" required:"false"`
	AwsRegion          string `envconfig:"AWS_REGION" required:"true"`
}

func ParseEnvConfiguration() (conf *AwsEnvConfig, err error) {
	conf = &AwsEnvConfig{}
	if err := envconfig.Process("", conf); err != nil {
		return nil, fmt.Errorf("Environment variables are not provided: %v ", err)
	}
	return
}
