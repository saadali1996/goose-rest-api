package main

import (
	"github.com/goose-rest-api/router"
	"github.com/rs/cors"
	"log"
	"net/http"
)

//go:generate swagger generate spec

// Package classification People API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:3000
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Jean-Thierry Bonhomme <jtbonhomme@gmail.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {

	// create router and start listen on port 8000
	router := router.NewRouter()
	log.Println("Serving APP")
	log.Fatal(http.ListenAndServe(":8000", setupGlobalMiddleware(router)))

}
