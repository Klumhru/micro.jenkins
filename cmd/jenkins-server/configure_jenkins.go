package main

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/IcelandairLabs/micro.jenkins/logic"
	"github.com/IcelandairLabs/micro.jenkins/restapi/jenkins"
	"github.com/IcelandairLabs/micro.jenkins/restapi/jenkins/health"
	"github.com/IcelandairLabs/micro.jenkins/restapi/jenkins/job"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureAPI(api *jenkins.JenkinsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

	api.JobGetJobListHandler = job.GetJobListHandlerFunc(func() middleware.Responder {
		return middleware.NotImplemented("operation job.GetJobList has not yet been implemented")
	})

	api.JobGetJobListHandler = job.GetJobListHandlerFunc(logic.GetJobListHandler)

	api.HealthGetServiceHealthHandler = health.GetServiceHealthHandlerFunc(func() middleware.Responder {
		return middleware.NotImplemented("operation health.GetServiceHealth has not yet been implemented")
	})

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
