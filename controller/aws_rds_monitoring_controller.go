package controller

import (
	"aws-rds-monitoring-tool/handler"
	"fmt"
)

type RdsMonitoringController struct {
	MonitoringHandler *handler.RdsMonitoringHandler
}

func NewRdsMonitoringController() (*RdsMonitoringController, error) {
	awsRdsClient, err := handler.NewAwsRdsClient()
	if err != nil {
		return nil, fmt.Errorf("Creating aws client failed: %v", err)
	}

	handler, err := handler.NewK8sMonitoringHandler(awsRdsClient)
	if err != nil {
		return nil, fmt.Errorf("Creating k8s monitoring handler failed: %v", err)
	}

	return &RdsMonitoringController{MonitoringHandler: handler}, nil
}
