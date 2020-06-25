package handler

import (
	"aws-rds-monitoring-tool/configuration"
	"aws-rds-monitoring-tool/models"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

type AwsRdsClient struct {
	awsRegion string
	rds       *rds.RDS
}

func NewAwsRdsClient(config *configuration.AwsEnvConfig) (*AwsRdsClient, error) {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region: &config.AwsRegion}))
	svc := rds.New(awsSession)

	return &AwsRdsClient{rds: svc, awsRegion: config.AwsRegion}, nil
}

func (kc *AwsRdsClient) GetAllRds(awsRegion string) (rdss models.RDSS, err error) {
	var rdsInstances *rds.DescribeDBInstancesOutput

	if awsRegion != kc.awsRegion {
		awsSession := session.Must(session.NewSession(&aws.Config{
			Region: &awsRegion}))
		rds := rds.New(awsSession)

		rdsInstances, err = rds.DescribeDBInstances(nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to get all rds instances: %v", err)
		}
	} else {
		rdsInstances, err = kc.rds.DescribeDBInstances(nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to get all rds instances: %v", err)
		}
	}

	rdsDbInstances := rdsInstances.DBInstances
	if len(rdsDbInstances) > 0 {
		rdss = make(models.RDSS, len(rdsDbInstances))
	}

	for _, rdsInstance := range rdsDbInstances {
		rds := &models.RDS{
			AvailabilityZone:   rdsInstance.AvailabilityZone,
			ClusterIdentifier:  rdsInstance.DBClusterIdentifier,
			InstanceIdentifier: rdsInstance.DBInstanceIdentifier,
			ResourceID:         rdsInstance.DbiResourceId,
			Status:             rdsInstance.DBInstanceStatus,
		}
		rdss = append(rdss, rds)
	}

	return
}
