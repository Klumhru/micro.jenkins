package jenkins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/spec"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/IcelandairLabs/micro.jenkins/restapi/jenkins/health"
	"github.com/IcelandairLabs/micro.jenkins/restapi/jenkins/job"
)

// NewJenkinsAPI creates a new Jenkins instance
func NewJenkinsAPI(spec *spec.Document) *JenkinsAPI {
	o := &JenkinsAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/com.icelandairlabs.jenkins.v2+json",
		defaultProduces: "application/com.icelandairlabs.jenkins.v2+json",
	}

	return o
}

/*JenkinsAPI A jenkins mirroring service for Dashboard use */
type JenkinsAPI struct {
	spec            *spec.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/com.icelandairlabs.jenkins.v2+json" mime type
	JSONConsumer httpkit.Consumer

	// JSONProducer registers a producer for a "application/com.icelandairlabs.jenkins.v2+json" mime type
	JSONProducer httpkit.Producer

	// JobGetJobListHandler sets the operation handler for the get job list operation
	JobGetJobListHandler job.GetJobListHandler
	// HealthGetServiceHealthHandler sets the operation handler for the get service health operation
	HealthGetServiceHealthHandler health.GetServiceHealthHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)
}

// SetDefaultProduces sets the default produces media type
func (o *JenkinsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *JenkinsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *JenkinsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *JenkinsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *JenkinsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *JenkinsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the JenkinsAPI
func (o *JenkinsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.JobGetJobListHandler == nil {
		unregistered = append(unregistered, "job.GetJobListHandler")
	}

	if o.HealthGetServiceHealthHandler == nil {
		unregistered = append(unregistered, "health.GetServiceHealthHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *JenkinsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *JenkinsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]httpkit.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *JenkinsAPI) ConsumersFor(mediaTypes []string) map[string]httpkit.Consumer {

	result := make(map[string]httpkit.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/com.icelandairlabs.jenkins.v2+json":
			result["application/com.icelandairlabs.jenkins.v2+json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *JenkinsAPI) ProducersFor(mediaTypes []string) map[string]httpkit.Producer {

	result := make(map[string]httpkit.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/com.icelandairlabs.jenkins.v2+json":
			result["application/com.icelandairlabs.jenkins.v2+json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *JenkinsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *JenkinsAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/jenkins/jobs"] = job.NewGetJobList(o.context, o.JobGetJobListHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/jenkins/health"] = health.NewGetServiceHealth(o.context, o.HealthGetServiceHealthHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *JenkinsAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
