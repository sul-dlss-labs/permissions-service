package main

import (
	"log"

	app "github.com/sul-dlss-labs/swagger-go-template"
	"github.com/sul-dlss-labs/swagger-go-template/config"
	"github.com/sul-dlss-labs/swagger-go-template/generated/restapi"
	"github.com/sul-dlss-labs/swagger-go-template/handlers"
)

func main() {
	rt, err := app.NewRuntime(config.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}

	server := createServer(rt)
	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func createServer(rt *app.Runtime) *restapi.Server {
	api := handlers.BuildAPI(rt)
	server := restapi.NewServer(api)
	server.SetHandler(handlers.BuildHandler(api))
	defer server.Shutdown()

	// set the port this service will be run on
	server.Port = rt.Config().Port
	return server
}
