package main

import (
	"log"

	permissions "github.com/sul-dlss-labs/permissions-service"
	"github.com/sul-dlss-labs/permissions-service/config"
	"github.com/sul-dlss-labs/permissions-service/generated/restapi"
	"github.com/sul-dlss-labs/permissions-service/handlers"
)

func main() {
	rt, err := permissions.NewRuntime(config.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}

	server := createServer(rt)
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func createServer(rt *permissions.Runtime) *restapi.Server {
	api := handlers.BuildAPI(rt)
	server := restapi.NewServer(api)
	server.SetHandler(handlers.BuildHandler(api))
	defer server.Shutdown()

	// set the port this service will be run on
	server.Port = rt.Config().Port
	return server
}
