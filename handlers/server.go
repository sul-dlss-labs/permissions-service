package handlers

import (
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/justinas/alice"
	app "github.com/sul-dlss-labs/permissions-service"
	"github.com/sul-dlss-labs/permissions-service/generated/restapi"
	"github.com/sul-dlss-labs/permissions-service/generated/restapi/operations"
	"github.com/sul-dlss-labs/permissions-service/middleware"
)

// BuildAPI create new service API
func BuildAPI(rt *app.Runtime) *operations.AppAPI {
	api := operations.NewAppAPI(swaggerSpec())
	// Add custom handlers here
	api.HealthCheckHandler = NewHealthCheck(rt)
	return api
}

// BuildHandler sets up the middleware that wraps the API
func BuildHandler(api *operations.AppAPI) http.Handler {
	return alice.New(
		middleware.NewHoneyBadgerMW(),
		middleware.NewRecoveryMW(),
		middleware.NewRequestLoggerMW(),
	).Then(api.Serve(nil))
}

func swaggerSpec() *loads.Document {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	return swaggerSpec
}
