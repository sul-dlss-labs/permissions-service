package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	permissions "github.com/sul-dlss-labs/permissions-service"
	"github.com/sul-dlss-labs/permissions-service/generated/models"
	"github.com/sul-dlss-labs/permissions-service/generated/restapi/operations"
)

// NewQueryAction will return the result of the query for authorization
func NewQueryAction(rt *permissions.Runtime) operations.QueryActionHandler {
	return &queryAction{}
}

type queryAction struct{}

// Handle the authorization query
func (d *queryAction) Handle(params operations.QueryActionParams, agent *permissions.Agent) middleware.Responder {
	return operations.NewQueryActionOK().WithPayload(&models.QueryResponse{Authorized: true})
}
