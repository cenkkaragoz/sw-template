package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexflint/go-arg"
	// "github.com/go-swagger/go-swagger/examples/cli/models"

	"github.com/cenkkaragoz/sw-template/swagger/models"
	"github.com/cenkkaragoz/sw-template/swagger/restapi"
	"github.com/cenkkaragoz/sw-template/swagger/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type cliArgs struct {
	Port int `arg:-p,help:port to listen to`
}

var (
	args = &cliArgs{
		Port: 8081,
	}
)

// func getHostnameHandler(params operations.GetHostnameParams) middleware.Responder {
// 	payload, err := os.Hostname()

// 	if err != nil {
// 		errPayload := &models.Error{
// 			Code:    500,
// 			Message: swag.String("failed to retrieve hostname"),
// 		}

// 		return operations.
// 			NewGetHostnameDefault(500).
// 			WithPayload(errPayload)
// 	}

// 	return operations.NewGetHostnameOK().WithPayload(payload)
// }

func main() {

	arg.MustParse(args)

	swaggerSpec, err := loads.Analyzed(restapi.FlatSwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)

	server.Port = args.Port
	defer server.Shutdown()

	api.GetHostnameHandler = operations.GetHostnameHandlerFunc(
		func(params operations.GetHostnameParams) middleware.Responder {
			response, err := os.Hostname()

			if err != nil {
				return operations.NewGetHostnameDefault(500).WithPayload(&models.Error{
					Code:    500,
					Message: swag.String("failed to retrieve hostname"),
				})
			}

			return operations.NewGetHostnameOK().WithPayload(response)
		})

	// Start listening using having the handlers and port
	// already set up.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	// Implement the handler functionality.
	// As all we need to do is give an implementation to the interface
	// we can just override the `api` method giving it a method with a valid
	// signature (we didn't need to have this

	fmt.Printf("port=%d\n", args.Port)
}
