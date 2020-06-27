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
GET /api/armt/v1/rdss?region=ap-south-1
```
Returns list of rds instances:
```
[
   {
          "availabilityZone": "ap-south-1b",
          "clusterIdentifier": null,
          "dbInstanceClass": "db.t2.micro",
          "dbName": sanjitdb,
          "engine": "mysql",
          "engineVersion": "8.0.16",
          "instanceIdentifier": "sanjit-database-1",
          "resourceId": "db-PX346I7MRIQVANIZ5XE6UB5YQY",
          "status": "available"
      },
       {
          "availabilityZone": "ap-east-1b",
          "clusterIdentifier": null,
          "dbInstanceClass": "db.t3.micro",
          "dbName": "sanjitdb2",
          "engine": "mysql",
          "engineVersion": "8.0.19",
          "instanceIdentifier": "sanjit-database-2",
          "resourceId": "db-Z5SH5B7G6MF3STWPY75S6RWGY4",
          "status": "creating"
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
      "query": "show tables;"
    }
  ],
  "dbName": "sanjitdb2",
  "dbEndpoint": "sanjit-database-2.rds.amazonaws.com"
}
```

## How to run
Update Makefile's "run" target and provide the following inputs -
```
* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY
* AWS_REGION
* RDS_USER_NAME
* RDS_PASSWORD
```

After you provided the above details, then you are good to run the application by simply executing -
```
make run
```

Now, you can make calls to the API exposed by this application - (for example)
```
curl -X GET "http://127.0.0.1:30001/api/armt/v1/rdss?region=ap-south-1"
```
## Technology Selection
* [Golang](https://golang.org/) used for implementation.
* [Go-Swagger](https://github.com/go-swagger/go-swagger) is used for rest service Swaggerization.
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) is used for talking to AWS cluster.
