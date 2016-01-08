package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetServiceHealthHandlerFunc turns a function with the right signature into a get service health handler
type GetServiceHealthHandlerFunc func() middleware.Responder

// Handle executing the request and returning a response
func (fn GetServiceHealthHandlerFunc) Handle() middleware.Responder {
	return fn()
}

// GetServiceHealthHandler interface for that can handle valid get service health params
type GetServiceHealthHandler interface {
	Handle() middleware.Responder
}

// NewGetServiceHealth creates a new http.Handler for the get service health operation
func NewGetServiceHealth(ctx *middleware.Context, handler GetServiceHealthHandler) *GetServiceHealth {
	return &GetServiceHealth{Context: ctx, Handler: handler}
}

/*GetServiceHealth swagger:route GET /api/jenkins/health health getServiceHealth

GetServiceHealth get service health API

*/
type GetServiceHealth struct {
	Context *middleware.Context
	Handler GetServiceHealthHandler
}

func (o *GetServiceHealth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, nil); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle() // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
