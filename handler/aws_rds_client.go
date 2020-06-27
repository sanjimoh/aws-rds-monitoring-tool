package handler

import (
	"aws-rds-monitoring-tool/configuration"
	"aws-rds-monitoring-tool/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/go-sql-driver/mysql"
	"log"
)

type AwsRdsClient struct {
	awsRegion   string
	rds         *rds.RDS
	rdsUser     string
	rdsPassword string
}

func NewAwsRdsClient(config *configuration.AwsEnvConfig) (*AwsRdsClient, error) {
	session := session.Must(session.NewSession(&aws.Config{
		Region: &config.AwsRegion}))
	svc := rds.New(session)

	return &AwsRdsClient{rds: svc, awsRegion: config.AwsRegion, rdsUser: config.RdsUser, rdsPassword: config.RdsPassword}, nil
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

func (kc *AwsRdsClient) ExecuteQueries(dbName string, dbEndpoint string, queries []string) error {
	dbConnectionStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=false", kc.rdsUser, kc.rdsPassword,
		dbEndpoint, dbName,
	)

	driver := mysql.MySQLDriver{}
	_ = driver

	db, err := sql.Open("mysql", dbConnectionStr)
	if err != nil {
		return fmt.Errorf("Failed to connect to db instance: %v", err)
	}

	for _, query := range queries {
 		rows, err := db.QueryContext(context.TODO(), query)
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

		if err := rows.Err(); err != nil {
			log.Printf("Error while iterating over rows: %v", err)
		}
	}

	return nil
}
