# aws-rds-monitoring-tool
## Scope
AWS RDS monitoring tool is meant for monitoring RDS resources in a cluster. It is still work in progress.

The existing codebase - (more to come..)
* Exposes a rest endpoint to it's client using which could return RDS instances including its current status in the 
  cluster.
  
## Rest interface definition and payload

```
GET /api/armt/v1/rdss
```
Returns list of rds instances:
```
[
  {
    "resourceId": "db-IXRXA2XS7KFFA6JWYYWFZEBJDE",
    "availabilityZone": "us-east-1b",
    "status" "available",
  },
  {
    "name": "db-YVS5NRBNHPGJZ3IT3WADXYSWYU",
    "availabilityZone": "us-east-1c",
    "status" "available",
  },
  ...
]
```


## Technology Selection
* [Golang](https://golang.org/) used for implementation.
* [Go-Swagger](https://github.com/go-swagger/go-swagger) is used for rest service Swaggerization.
* [aws-sdk-go](https://github.com/aws/aws-sdk-go) is used for talking to AWS cluster.
