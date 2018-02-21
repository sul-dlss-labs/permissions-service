// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Template Application",
    "title": "app",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "0.1.0"
  },
  "host": "app.dlss.stanford.edu",
  "basePath": "/v1",
  "paths": {
    "/healthcheck": {
      "get": {
        "description": "The healthcheck endpoint provides information about the health of the service.",
        "summary": "Health Check",
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "The service is functioning nominally",
            "schema": {
              "$ref": "#/definitions/HealthCheckResponse"
            }
          },
          "503": {
            "description": "The service is not working correctly",
            "schema": {
              "$ref": "#/definitions/HealthCheckResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "detail": {
          "description": "a human-readable explanation specific to this occurrence of the problem.",
          "type": "string",
          "example": "Title must contain at least three characters."
        },
        "source": {
          "type": "object",
          "properties": {
            "pointer": {
              "type": "string",
              "example": "/data/attributes/title"
            }
          }
        },
        "title": {
          "description": "a short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem.",
          "type": "string",
          "example": "Invalid Attribute"
        }
      }
    },
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "errors": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Error"
          }
        }
      }
    },
    "HealthCheckResponse": {
      "type": "object",
      "properties": {
        "status": {
          "description": "The status of the service",
          "type": "string"
        }
      },
      "example": {
        "status": "OK"
      }
    }
  }
}`))
}
