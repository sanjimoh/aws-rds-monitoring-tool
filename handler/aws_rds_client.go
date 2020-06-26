package handler

import (
	"aws-rds-monitoring-tool/configuration"
	"aws-rds-monitoring-tool/models"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	"log"
)

type AwsRdsClient struct {
	awsRegion  string
	awsSession *session.Session
	rds        *rds.RDS
}

func NewAwsRdsClient(config *configuration.AwsEnvConfig) (*AwsRdsClient, error) {
	session := session.Must(session.NewSession(&aws.Config{
		Region: &config.AwsRegion}))
	svc := rds.New(session)

	return &AwsRdsClient{rds: svc, awsRegion: config.AwsRegion, awsSession: session}, nil
}

func (kc *AwsRdsClient) GetAllRds(awsRegion string) (rdss models.RDSS, err error) {
	var rdsInstances *rds.DescribeDBInstancesOutput

	if len(awsRegion) == 0 || awsRegion == kc.awsRegion {
		rdsInstances, err = kc.rds.DescribeDBInstances(nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to get all rds instances: %v", err)
		}
	} else if awsRegion != kc.awsRegion {
		awsSession := session.Must(session.NewSession(&aws.Config{
			Region: &awsRegion}))
		rds := rds.New(awsSession)

		rdsInstances, err = rds.DescribeDBInstances(nil)
		if err != nil {
			return nil, fmt.Errorf("Failed to get all rds instances: %v", err)
		}
	}

	rdsDbInstances := rdsInstances.DBInstances
	for _, rdsInstance := range rdsDbInstances {
		rds := &models.RDS{
			AvailabilityZone:   rdsInstance.AvailabilityZone,
			ClusterIdentifier:  rdsInstance.DBClusterIdentifier,
			DbInstanceClass:    rdsInstance.DBInstanceClass,
			DbName:             rdsInstance.DBName,
			Engine:             rdsInstance.Engine,
			EngineVersion:      rdsInstance.EngineVersion,
			InstanceIdentifier: rdsInstance.DBInstanceIdentifier,
			ResourceID:         rdsInstance.DbiResourceId,
			Status:             rdsInstance.DBInstanceStatus,
		}
		rdss = append(rdss, rds)
	}

	return
}

func (kc *AwsRdsClient) ExecuteQueries(region string, dbUser string, dbName string, dbEndpoint string, iamArn string, queries []string) error {
	session := session.Must(session.NewSession(&aws.Config{Region: &region}))
	credentials := stscreds.NewCredentials(session, iamArn)

	token, err := rdsutils.BuildAuthToken(dbEndpoint, region, dbUser, credentials)
	if err != nil {
		return fmt.Errorf("Failed to get auth token for the rds instance: %v", err)
	}

	dbConnectionStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true",
		dbUser, token, dbEndpoint, dbName,
	)

	db, err := sql.Open("mysql", dbConnectionStr)
	if err != nil {
		return fmt.Errorf("Failed to connect to db instance: %v", err)
	}

	for _, query := range queries {
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("Query %s exection failed: %v", query, err)
			continue
		}

		for rows.Next() {
			data := make(map[string]string)

			cols, err := rows.Columns()
			if err != nil {
				log.Printf("failed to read columns from row: %v", err)
				continue
			}

			columns := make([]string, len(cols))
			columnPtr := make([]interface{}, len(cols))
			for i, _ := range columns {
				columnPtr[i] = &columns[i]
			}

			err = rows.Scan(columnPtr...)
			if err != nil {
				log.Printf("failed to read row: %v", err)
				continue
			}

			for i, colName := range cols {
				data[colName] = columns[i]
			}

			log.Printf("Retrieved row : %v", data)
		}
	}

	return nil
}
