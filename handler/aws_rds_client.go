package handler

import (
	"aws-rds-monitoring-tool/models"
)

type AwsRdsClient struct{}

func NewAwsRdsClient() (*AwsRdsClient, error) {
	return &AwsRdsClient{}, nil
}

func (kc *AwsRdsClient) GetAllRds() (rdss models.RDSS, err error) {
	return
}
