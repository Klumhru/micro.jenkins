package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/IcelandairLabs/micro.jenkins/models"
)

/*GetServiceHealthOK OK

swagger:response getServiceHealthOK
*/
type GetServiceHealthOK struct {

	// In: body
	Payload *models.Health `json:"body,omitempty"`
}

// NewGetServiceHealthOK creates GetServiceHealthOK with default headers values
func NewGetServiceHealthOK() *GetServiceHealthOK {
	return &GetServiceHealthOK{}
}

// WithPayload adds the payload to the get service health o k response
func (o *GetServiceHealthOK) WithPayload(payload *models.Health) *GetServiceHealthOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *GetServiceHealthOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
