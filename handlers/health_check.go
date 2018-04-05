package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	permissions "github.com/sul-dlss-labs/permissions-service"
	"github.com/sul-dlss-labs/permissions-service/generated/models"
	"github.com/sul-dlss-labs/permissions-service/generated/restapi/operations"
)

// NewHealthCheck will return the service health
func NewHealthCheck(rt *permissions.Runtime) operations.HealthCheckHandler {
	return &healthCheck{}
}

type healthCheck struct{}

// Handle the health check request
func (d *healthCheck) Handle(params operations.HealthCheckParams) middleware.Responder {
	return operations.NewHealthCheckOK().WithPayload(&models.HealthCheckResponse{Status: "OK"})
}
