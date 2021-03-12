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
  "swagger": "2.0",
  "info": {
    "title": "Goldpinger",
    "version": "3.0.0"
  },
  "paths": {
    "/check": {
      "get": {
        "description": "Queries the API server for all other pods in this service, and pings them via their pods IPs. Calls their /ping endpoint",
        "produces": [
          "application/json"
        ],
        "operationId": "checkServicePods",
        "responses": {
          "200": {
            "description": "Success, return response",
            "schema": {
              "$ref": "#/definitions/CheckResults"
            }
          }
        }
      }
    },
    "/check_all": {
      "get": {
        "description": "Queries the API server for all other pods in this service, and makes all of them query all of their neighbours, using their pods IPs. Calls their /check endpoint.",
        "produces": [
          "application/json"
        ],
        "operationId": "checkAllPods",
        "responses": {
          "200": {
            "description": "Success, return response",
            "schema": {
              "$ref": "#/definitions/CheckAllResults"
            }
          }
        }
      }
    },
    "/cluster_health": {
      "get": {
        "description": "Checks the full graph. Returns a binary OK or not OK.",
        "produces": [
          "application/json"
        ],
        "operationId": "clusterHealth",
        "responses": {
          "200": {
            "description": "Healthy cluster",
            "schema": {
              "$ref": "#/definitions/ClusterHealthResults"
            }
          },
          "503": {
            "description": "Unhealthy cluster",
            "schema": {
              "$ref": "#/definitions/ClusterHealthResults"
            }
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "description": "The healthcheck endpoint provides detailed information about the health of a web service. If each of the components required by the service are healthy, then the service is considered healthy and will return a 200 OK response. If any of the components needed by the service are unhealthy, then a 503 Service Unavailable response will be provided.",
        "produces": [
          "application/json"
        ],
        "operationId": "healthz",
        "responses": {
          "200": {
            "description": "Health check report",
            "schema": {
              "$ref": "#/definitions/HealthCheckResults"
            }
          },
          "503": {
            "description": "Unhealthy service",
            "schema": {
              "$ref": "#/definitions/HealthCheckResults"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "description": "return query stats",
        "produces": [
          "application/json"
        ],
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "return success",
            "schema": {
              "$ref": "#/definitions/PingResults"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CallStats": {
      "properties": {
        "check": {
          "type": "integer"
        },
        "check_all": {
          "type": "integer"
        },
        "ping": {
          "type": "integer"
        }
      }
    },
    "CheckAllPodResult": {
      "type": "object",
      "properties": {
        "HostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "PodIP": {
          "type": "string",
          "format": "ipv4"
        },
        "error": {
          "type": "string"
        },
        "response": {
          "$ref": "#/definitions/CheckResults"
        },
        "status-code": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "CheckAllResults": {
      "type": "object",
      "properties": {
        "OK": {
          "type": "boolean",
          "default": false
        },
        "dnsResults": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/DnsResults"
          }
        },
        "hosts": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "hostIP": {
                "type": "string",
                "format": "ipv4"
              },
              "podIP": {
                "type": "string",
                "format": "ipv4"
              },
              "podName": {
                "type": "string"
              }
            }
          }
        },
        "hosts-healthy": {
          "type": "integer",
          "format": "int32"
        },
        "hosts-number": {
          "type": "integer",
          "format": "int32"
        },
        "responses": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/CheckAllPodResult"
          }
        }
      }
    },
    "CheckResults": {
      "type": "object",
      "properties": {
        "dnsResults": {
          "$ref": "#/definitions/DnsResults"
        },
        "podResults": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/PodResult"
          }
        }
      }
    },
    "ClusterHealthResults": {
      "type": "object",
      "required": [
        "OK"
      ],
      "properties": {
        "DNS": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "duration-ns": {
          "type": "integer",
          "format": "int64"
        },
        "generated-at": {
          "type": "string",
          "format": "date-time"
        },
        "hostsHealthy": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "hostsUnhealthy": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "podsHealthy": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "podsUnhealthy": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "DnsResult": {
      "properties": {
        "error": {
          "type": "string"
        },
        "response-time-ms": {
          "type": "number",
          "format": "int64"
        }
      }
    },
    "DnsResults": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/DnsResult"
      }
    },
    "HealthCheckResults": {
      "type": "object",
      "properties": {
        "OK": {
          "type": "boolean",
          "default": false
        },
        "duration-ns": {
          "type": "integer",
          "format": "int64"
        },
        "generated-at": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "PingResults": {
      "type": "object",
      "properties": {
        "boot_time": {
          "type": "string",
          "format": "date-time"
        },
        "received": {
          "$ref": "#/definitions/CallStats"
        }
      }
    },
    "PodResult": {
      "type": "object",
      "properties": {
        "HostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "PingTime": {
          "type": "string",
          "format": "date-time"
        },
        "PodIP": {
          "type": "string",
          "format": "ipv4"
        },
        "error": {
          "type": "string"
        },
        "response": {
          "$ref": "#/definitions/PingResults"
        },
        "response-time-ms": {
          "description": "wall clock time in milliseconds",
          "type": "number",
          "format": "int64"
        },
        "status-code": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Goldpinger",
    "version": "3.0.0"
  },
  "paths": {
    "/check": {
      "get": {
        "description": "Queries the API server for all other pods in this service, and pings them via their pods IPs. Calls their /ping endpoint",
        "produces": [
          "application/json"
        ],
        "operationId": "checkServicePods",
        "responses": {
          "200": {
            "description": "Success, return response",
            "schema": {
              "$ref": "#/definitions/CheckResults"
            }
          }
        }
      }
    },
    "/check_all": {
      "get": {
        "description": "Queries the API server for all other pods in this service, and makes all of them query all of their neighbours, using their pods IPs. Calls their /check endpoint.",
        "produces": [
          "application/json"
        ],
        "operationId": "checkAllPods",
        "responses": {
          "200": {
            "description": "Success, return response",
            "schema": {
              "$ref": "#/definitions/CheckAllResults"
            }
          }
        }
      }
    },
    "/cluster_health": {
      "get": {
        "description": "Checks the full graph. Returns a binary OK or not OK.",
        "produces": [
          "application/json"
        ],
        "operationId": "clusterHealth",
        "responses": {
          "200": {
            "description": "Healthy cluster",
            "schema": {
              "$ref": "#/definitions/ClusterHealthResults"
            }
          },
          "503": {
            "description": "Unhealthy cluster",
            "schema": {
              "$ref": "#/definitions/ClusterHealthResults"
            }
          }
        }
      }
    },
    "/healthz": {
      "get": {
        "description": "The healthcheck endpoint provides detailed information about the health of a web service. If each of the components required by the service are healthy, then the service is considered healthy and will return a 200 OK response. If any of the components needed by the service are unhealthy, then a 503 Service Unavailable response will be provided.",
        "produces": [
          "application/json"
        ],
        "operationId": "healthz",
        "responses": {
          "200": {
            "description": "Health check report",
            "schema": {
              "$ref": "#/definitions/HealthCheckResults"
            }
          },
          "503": {
            "description": "Unhealthy service",
            "schema": {
              "$ref": "#/definitions/HealthCheckResults"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "description": "return query stats",
        "produces": [
          "application/json"
        ],
        "operationId": "ping",
        "responses": {
          "200": {
            "description": "return success",
            "schema": {
              "$ref": "#/definitions/PingResults"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CallStats": {
      "properties": {
        "check": {
          "type": "integer"
        },
        "check_all": {
          "type": "integer"
        },
        "ping": {
          "type": "integer"
        }
      }
    },
    "CheckAllPodResult": {
      "type": "object",
      "properties": {
        "HostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "PodIP": {
          "type": "string",
          "format": "ipv4"
        },
        "error": {
          "type": "string"
        },
        "response": {
          "$ref": "#/definitions/CheckResults"
        },
        "status-code": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "CheckAllResults": {
      "type": "object",
      "properties": {
        "OK": {
          "type": "boolean",
          "default": false
        },
        "dnsResults": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/DnsResults"
          }
        },
        "hosts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/CheckAllResultsHostsItems0"
          }
        },
        "hosts-healthy": {
          "type": "integer",
          "format": "int32"
        },
        "hosts-number": {
          "type": "integer",
          "format": "int32"
        },
        "responses": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/CheckAllPodResult"
          }
        }
      }
    },
    "CheckAllResultsHostsItems0": {
      "type": "object",
      "properties": {
        "hostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "podIP": {
          "type": "string",
          "format": "ipv4"
        },
        "podName": {
          "type": "string"
        }
      }
    },
    "CheckResults": {
      "type": "object",
      "properties": {
        "dnsResults": {
          "$ref": "#/definitions/DnsResults"
        },
        "podResults": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/PodResult"
          }
        }
      }
    },
    "ClusterHealthResults": {
      "type": "object",
      "required": [
        "OK"
      ],
      "properties": {
        "DNS": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "duration-ns": {
          "type": "integer",
          "format": "int64"
        },
        "generated-at": {
          "type": "string",
          "format": "date-time"
        },
        "hostsHealthy": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "hostsUnhealthy": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "podsHealthy": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "podsUnhealthy": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "DnsResult": {
      "properties": {
        "error": {
          "type": "string"
        },
        "response-time-ms": {
          "type": "number",
          "format": "int64"
        }
      }
    },
    "DnsResults": {
      "type": "object",
      "additionalProperties": {
        "$ref": "#/definitions/DnsResult"
      }
    },
    "HealthCheckResults": {
      "type": "object",
      "properties": {
        "OK": {
          "type": "boolean",
          "default": false
        },
        "duration-ns": {
          "type": "integer",
          "format": "int64"
        },
        "generated-at": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "PingResults": {
      "type": "object",
      "properties": {
        "boot_time": {
          "type": "string",
          "format": "date-time"
        },
        "received": {
          "$ref": "#/definitions/CallStats"
        }
      }
    },
    "PodResult": {
      "type": "object",
      "properties": {
        "HostIP": {
          "type": "string",
          "format": "ipv4"
        },
        "OK": {
          "type": "boolean",
          "default": false
        },
        "PingTime": {
          "type": "string",
          "format": "date-time"
        },
        "PodIP": {
          "type": "string",
          "format": "ipv4"
        },
        "error": {
          "type": "string"
        },
        "response": {
          "$ref": "#/definitions/PingResults"
        },
        "response-time-ms": {
          "description": "wall clock time in milliseconds",
          "type": "number",
          "format": "int64"
        },
        "status-code": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}`))
}
