package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"taskmanager/internal/generated/restapi"
	"taskmanager/internal/generated/restapi/operations"
	handlers "taskmanager/api"
)

func main() {
	var portFlag = flag.Int("port", 8080, "Port to run this service on")
	flag.Parse()

	// Load the OpenAPI spec
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Create new service API
	api := operations.NewTaskmanagerAPI(swaggerSpec)
	
	// Register API handlers
	handlers.RegisterHandlers(api)
	
	// Configure API
	server := restapi.NewServer(api)
	server.Port = *portFlag
	
	// Start server
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
