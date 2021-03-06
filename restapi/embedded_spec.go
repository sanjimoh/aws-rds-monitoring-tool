// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "ARMT provides rest endpoints to help clients to get monitoring data from RDS instances in AWS cluster.",
    "title": "AWS RDS Monitoring Tool (ARMT)",
    "version": "1.0.0"
  },
  "basePath": "/api/armt",
  "paths": {
    "/v1/rds/queries": {
      "post": {
        "description": "For example:\n` + "`" + `` + "`" + `` + "`" + `\nPOST /api/armt/v1/rds/queries\n` + "`" + `` + "`" + `` + "`" + `\nSample request body will be:\n` + "`" + `` + "`" + `` + "`" + `\n{\n  \"queries\": [\n    {\n      \"query\": \"string\"\n    }\n  ],\n  \"dbName\": \"string\",\n  \"dbEndpoint\": \"string\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "aws-rds-monitoring-tool"
        ],
        "summary": "Execute a set of queries on a RDS instance",
        "parameters": [
          {
            "description": "RDS query parameter.",
            "name": "RdsQueriesExecAttr",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/RdsQueriesExecAttr"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "RDS queries execution successfull.",
            "schema": {
              "$ref": "#/definitions/RdsQueriesExecAttr"
            }
          },
          "400": {
            "description": "Bad Request, Unable to execute RDS queries.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/rdss": {
      "get": {
        "description": "For example:\n` + "`" + `` + "`" + `` + "`" + `\nGET /api/armt/v1/rdss?region=us-east-1b\n` + "`" + `` + "`" + `` + "`" + `\n\nReturns list of rds instances:\n` + "`" + `` + "`" + `` + "`" + `\n[\n  {\n            \"availabilityZone\": \"ap-south-1b\",\n            \"clusterIdentifier\": null,\n            \"dbInstanceClass\": \"db.t2.micro\",\n            \"dbName\": sanjitdb,\n            \"engine\": \"mysql\",\n            \"engineVersion\": \"8.0.16\",\n            \"instanceIdentifier\": \"sanjit-database-1\",\n            \"resourceId\": \"db-PX346I7MRIQVANIZ5XE6UB5YQY\",\n            \"status\": \"available\"\n        },\n        {\n          \"availabilityZone\": \"ap-east-1b\",\n          \"clusterIdentifier\": null,\n          \"dbInstanceClass\": \"db.t3.micro\",\n          \"dbName\": \"sanjitdb2\",\n          \"engine\": \"mysql\",\n          \"engineVersion\": \"8.0.19\",\n          \"instanceIdentifier\": \"sanjit-database-2\",\n          \"resourceId\": \"db-Z5SH5B7G6MF3STWPY75S6RWGY4\",\n          \"status\": \"creating\"\n        },\n        ...\n]\n` + "`" + `` + "`" + `` + "`" + `\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "aws-rds-monitoring-tool"
        ],
        "summary": "Gets all rds statistics running in the AWS cluster",
        "parameters": [
          {
            "type": "string",
            "description": "Pass an aws region for which RDS instance status is required to be monitored. If region is not passed, then the default AWS region passed in env is taken.",
            "name": "region",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Fetching of all pod status is successful.",
            "schema": {
              "$ref": "#/definitions/RDSS"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "RDS": {
      "description": "An RDS instance",
      "type": "object",
      "required": [
        "resourceId",
        "clusterIdentifier",
        "instanceIdentifier",
        "availabilityZone",
        "dbInstanceClass",
        "dbName",
        "engine",
        "engineVersion",
        "status"
      ],
      "properties": {
        "availabilityZone": {
          "description": "Specifies the name of the Availability Zone the DB instance is located in.",
          "type": "string"
        },
        "clusterIdentifier": {
          "description": "If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.",
          "type": "string"
        },
        "dbInstanceClass": {
          "description": "Contains the name of the compute and memory capacity class of the DB instance.",
          "type": "string"
        },
        "dbName": {
          "description": "Contains the name of the initial database of this instance that was provided at the time of create.",
          "type": "string"
        },
        "engine": {
          "description": "Provides the name of the database engine to be used for this DB instance.",
          "type": "string"
        },
        "engineVersion": {
          "description": "Indicates the database engine version.",
          "type": "string"
        },
        "instanceIdentifier": {
          "description": "Contains a user-supplied database identifier. This identifier is the unique key that identifies a DB instance.",
          "type": "string"
        },
        "resourceId": {
          "description": "The AWS Region-unique, immutable identifier for the DB instance.",
          "type": "string"
        },
        "status": {
          "description": "The status of a DB instance indicates the health of the DB instance. More refer here - https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Status.html",
          "type": "string"
        }
      }
    },
    "RDSQuery": {
      "description": "An RDS query",
      "type": "object",
      "required": [
        "query"
      ],
      "properties": {
        "query": {
          "description": "An RDS query",
          "type": "string"
        }
      }
    },
    "RDSS": {
      "description": "Array of RDS items",
      "type": "array",
      "items": {
        "$ref": "#/definitions/RDS"
      }
    },
    "RdsQueriesExecAttr": {
      "description": "Attributes required for RDS queries executions.",
      "type": "object",
      "required": [
        "dbName",
        "dbEndpoint",
        "queries"
      ],
      "properties": {
        "dbEndpoint": {
          "description": "DB endpoint",
          "type": "string"
        },
        "dbName": {
          "description": "DB Name",
          "type": "string"
        },
        "queries": {
          "description": "Array of query",
          "type": "array",
          "items": {
            "$ref": "#/definitions/RDSQuery"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "ARMT provides rest endpoints to help clients to get monitoring data from RDS instances in AWS cluster.",
    "title": "AWS RDS Monitoring Tool (ARMT)",
    "version": "1.0.0"
  },
  "basePath": "/api/armt",
  "paths": {
    "/v1/rds/queries": {
      "post": {
        "description": "For example:\n` + "`" + `` + "`" + `` + "`" + `\nPOST /api/armt/v1/rds/queries\n` + "`" + `` + "`" + `` + "`" + `\nSample request body will be:\n` + "`" + `` + "`" + `` + "`" + `\n{\n  \"queries\": [\n    {\n      \"query\": \"string\"\n    }\n  ],\n  \"dbName\": \"string\",\n  \"dbEndpoint\": \"string\"\n}\n` + "`" + `` + "`" + `` + "`" + `\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "aws-rds-monitoring-tool"
        ],
        "summary": "Execute a set of queries on a RDS instance",
        "parameters": [
          {
            "description": "RDS query parameter.",
            "name": "RdsQueriesExecAttr",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/RdsQueriesExecAttr"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "RDS queries execution successfull.",
            "schema": {
              "$ref": "#/definitions/RdsQueriesExecAttr"
            }
          },
          "400": {
            "description": "Bad Request, Unable to execute RDS queries.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/v1/rdss": {
      "get": {
        "description": "For example:\n` + "`" + `` + "`" + `` + "`" + `\nGET /api/armt/v1/rdss?region=us-east-1b\n` + "`" + `` + "`" + `` + "`" + `\n\nReturns list of rds instances:\n` + "`" + `` + "`" + `` + "`" + `\n[\n  {\n            \"availabilityZone\": \"ap-south-1b\",\n            \"clusterIdentifier\": null,\n            \"dbInstanceClass\": \"db.t2.micro\",\n            \"dbName\": sanjitdb,\n            \"engine\": \"mysql\",\n            \"engineVersion\": \"8.0.16\",\n            \"instanceIdentifier\": \"sanjit-database-1\",\n            \"resourceId\": \"db-PX346I7MRIQVANIZ5XE6UB5YQY\",\n            \"status\": \"available\"\n        },\n        {\n          \"availabilityZone\": \"ap-east-1b\",\n          \"clusterIdentifier\": null,\n          \"dbInstanceClass\": \"db.t3.micro\",\n          \"dbName\": \"sanjitdb2\",\n          \"engine\": \"mysql\",\n          \"engineVersion\": \"8.0.19\",\n          \"instanceIdentifier\": \"sanjit-database-2\",\n          \"resourceId\": \"db-Z5SH5B7G6MF3STWPY75S6RWGY4\",\n          \"status\": \"creating\"\n        },\n        ...\n]\n` + "`" + `` + "`" + `` + "`" + `\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "aws-rds-monitoring-tool"
        ],
        "summary": "Gets all rds statistics running in the AWS cluster",
        "parameters": [
          {
            "type": "string",
            "description": "Pass an aws region for which RDS instance status is required to be monitored. If region is not passed, then the default AWS region passed in env is taken.",
            "name": "region",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Fetching of all pod status is successful.",
            "schema": {
              "$ref": "#/definitions/RDSS"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "RDS": {
      "description": "An RDS instance",
      "type": "object",
      "required": [
        "resourceId",
        "clusterIdentifier",
        "instanceIdentifier",
        "availabilityZone",
        "dbInstanceClass",
        "dbName",
        "engine",
        "engineVersion",
        "status"
      ],
      "properties": {
        "availabilityZone": {
          "description": "Specifies the name of the Availability Zone the DB instance is located in.",
          "type": "string"
        },
        "clusterIdentifier": {
          "description": "If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.",
          "type": "string"
        },
        "dbInstanceClass": {
          "description": "Contains the name of the compute and memory capacity class of the DB instance.",
          "type": "string"
        },
        "dbName": {
          "description": "Contains the name of the initial database of this instance that was provided at the time of create.",
          "type": "string"
        },
        "engine": {
          "description": "Provides the name of the database engine to be used for this DB instance.",
          "type": "string"
        },
        "engineVersion": {
          "description": "Indicates the database engine version.",
          "type": "string"
        },
        "instanceIdentifier": {
          "description": "Contains a user-supplied database identifier. This identifier is the unique key that identifies a DB instance.",
          "type": "string"
        },
        "resourceId": {
          "description": "The AWS Region-unique, immutable identifier for the DB instance.",
          "type": "string"
        },
        "status": {
          "description": "The status of a DB instance indicates the health of the DB instance. More refer here - https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Status.html",
          "type": "string"
        }
      }
    },
    "RDSQuery": {
      "description": "An RDS query",
      "type": "object",
      "required": [
        "query"
      ],
      "properties": {
        "query": {
          "description": "An RDS query",
          "type": "string"
        }
      }
    },
    "RDSS": {
      "description": "Array of RDS items",
      "type": "array",
      "items": {
        "$ref": "#/definitions/RDS"
      }
    },
    "RdsQueriesExecAttr": {
      "description": "Attributes required for RDS queries executions.",
      "type": "object",
      "required": [
        "dbName",
        "dbEndpoint",
        "queries"
      ],
      "properties": {
        "dbEndpoint": {
          "description": "DB endpoint",
          "type": "string"
        },
        "dbName": {
          "description": "DB Name",
          "type": "string"
        },
        "queries": {
          "description": "Array of query",
          "type": "array",
          "items": {
            "$ref": "#/definitions/RDSQuery"
          }
        }
      }
    }
  }
}`))
}
