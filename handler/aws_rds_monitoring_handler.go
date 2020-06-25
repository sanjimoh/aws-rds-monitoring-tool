package handler

import (
	"aws-rds-monitoring-tool/models"
	"fmt"
)

type RdsMonitoringHandler struct {
	awsRdsClient *AwsRdsClient
}

func NewK8sMonitoringHandler(awsClient *AwsRdsClient) (*RdsMonitoringHandler, error) {
	return &RdsMonitoringHandler{awsRdsClient: awsClient}, nil
}

func (rmh *RdsMonitoringHandler) GetV1Rdss(awsRegion string) (rdss models.RDSS, err error) {
	rdss, err = rmh.awsRdsClient.GetAllRds(awsRegion)
	if err != nil {
		return nil, fmt.Errorf("Fetching rds instances failed: %v", err)
	}

	return
}

func (rmh *RdsMonitoringHandler) PostV1RdsQueries(queryExecAttrs *models.RdsQueriesExecAttr) error {
	queries := make([]string, len(queryExecAttrs.Queries))
	for _, query := range queryExecAttrs.Queries {
		queries = append(queries, *query.Query)
	}

	return rmh.awsRdsClient.ExecuteQueries(*queryExecAttrs.Region, *queryExecAttrs.DbUser, *queryExecAttrs.DbName, *queryExecAttrs.DbEndpoint, *queryExecAttrs.IamArn, queries)
}
