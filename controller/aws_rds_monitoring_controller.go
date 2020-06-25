package controller

import (
	"aws-rds-monitoring-tool/configuration"
	"aws-rds-monitoring-tool/handler"
	"fmt"
)

type RdsMonitoringController struct {
	MonitoringHandler *handler.RdsMonitoringHandler
}

func NewRdsMonitoringController() (*RdsMonitoringController, error) {
	config, err := configuration.ParseEnvConfiguration()
	if err != nil {
		return nil, fmt.Errorf("Could not parse aws rds monitoring tool service config: %s", err)
	}

	awsRdsClient, err := handler.NewAwsRdsClient(config)
	if err != nil {
		return nil, fmt.Errorf("Creating aws client failed: %v", err)
	}

	handler, err := handler.NewK8sMonitoringHandler(awsRdsClient)
	if err != nil {
		return nil, fmt.Errorf("Creating k8s monitoring handler failed: %v", err)
	}

	return &RdsMonitoringController{MonitoringHandler: handler}, nil
}
