package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	app "github.com/sul-dlss-labs/swagger-go-template"
	"github.com/sul-dlss-labs/swagger-go-template/generated/models"
	"github.com/sul-dlss-labs/swagger-go-template/generated/restapi/operations"
)

// NewHealthCheck will return the service health
func NewHealthCheck(rt *app.Runtime) operations.HealthCheckHandler {
	return &healthCheck{}
}

type healthCheck struct{}

// Handle the health check request
func (d *healthCheck) Handle(params operations.HealthCheckParams) middleware.Responder {
	return operations.NewHealthCheckOK().WithPayload(&models.HealthCheckResponse{Status: "OK"})
}
