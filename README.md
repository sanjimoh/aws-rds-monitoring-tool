# aws-rds-monitoring-tool
## Scope
AWS RDS monitoring tool is meant for monitoring RDS resources in a cluster. It is still work in progress.

The existing codebase - (more to come..)
* Exposes a rest endpoint to it's client using which could return RDS instances including its current status in the 
  cluster. If region is not passed, then the default AWS region passed in env is taken.
* Now there is support for execution of queries on RDS instances. 

## Rest interface definition and payload
* Get all RDS instances and its status:
```
GET /api/armt/v1/rdss?region=us-east-1b
```
Returns list of rds instances:
```
[
  {
    "resourceId": "db-IXRXA2XS7KFFA6JWYYWFZEBJDE",
    "clusterIdentifier": "",
    "instanceIdentifier": "mysqldb",
    "availabilityZone": "us-east-1b",
    "status" "available",
  },
  {
    "resourceId": "db-YVS5NRBNHPGJZ3IT3WADXYSWYU",
    "clusterIdentifier": "",
    "instanceIdentifier": "mysqldb",
    "availabilityZone": "us-east-1b",
    "status" "backing-up",
  },
  ...
]
```
* Execute queries on RDS instance:
```
POST /api/armt/v1/rds/queries
```
Sample request body will be:
```
{
  "queries": [
    {
      "query": "string"
    }
  ],
  "region": "string",
  "dbUser": "string",
  "dbName": "string",
  "dbEndpoint": "string",
  "iamArn": "string"
}
```


## Technology Selection
* [Golang](https://golang.org/) used for implementation.
* [Go-Swagger](https://github.com/go-swagger/go-swagger) is used for rest service Swaggerization.
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) is used for talking to AWS cluster.
