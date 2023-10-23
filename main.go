package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/oyeprashar/microservice-boilerplate/internal/router/v1"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func main() {

	/*
		TODO: Logging the request/response
			1. We need to pass the logger all over the code base instead of creating multiple logger object
			2. We can then use in the the handler to log the body
	*/

	/*
		TODO:
			1. Complete the basic things for the service
			2. Add features
			3. Dockerise
	*/

	/*
		TODO: Things to add
			1. Handler with all the required middleware
			2. The service should log all the request and response that are being processes
			3. Prometheus and Grafana
			4. MongoDB for logging the request and response
			5. Unit test for this boiler plate (in the most organized way, probably in a single folder (Check))
			6. Add configs based on the ENVs. Local | Dev | Prod
	*/

	// TODO: This file should be as minimal as possible. So refactor it
	logrusLogger := logrus.New()
	logrusLogger.SetOutput(os.Stdout)

	/*
	 Chi is a lightweight, fast, and flexible HTTP router for building web applications in Go.
	 Routers in Go are used to route incoming HTTP requests to the appropriate handlers or endpoints
	 based on the request's path and method.
	*/

	r := chi.NewRouter()
	r.Use(middleware.RequestID) // This middleware generates a unique identifier (usually a UUID or GUID)
	//r.Use(middleware.Logger)
	// TODO: add a middleware that prints the request and the response of the API made to us

	r.Route("/v1", v1.V1Router)

	// understand this code very clearly
	server := &http.Server{
		Addr:              ":3333",
		ReadHeaderTimeout: 3 * time.Minute,
		Handler:           r, // TODO: we need to define this handler
	}

	logrusLogger.Info("-- SERVER STARTED --")
	logrusLogger.Info("Listening at port", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
