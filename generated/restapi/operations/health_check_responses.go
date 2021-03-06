// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sul-dlss-labs/permissions-service/generated/models"
)

// HealthCheckOKCode is the HTTP code returned for type HealthCheckOK
const HealthCheckOKCode int = 200

/*HealthCheckOK The service is functioning nominally

swagger:response healthCheckOK
*/
type HealthCheckOK struct {

	/*
	  In: Body
	*/
	Payload *models.HealthCheckResponse `json:"body,omitempty"`
}

// NewHealthCheckOK creates HealthCheckOK with default headers values
func NewHealthCheckOK() *HealthCheckOK {
	return &HealthCheckOK{}
}

// WithPayload adds the payload to the health check o k response
func (o *HealthCheckOK) WithPayload(payload *models.HealthCheckResponse) *HealthCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the health check o k response
func (o *HealthCheckOK) SetPayload(payload *models.HealthCheckResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HealthCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HealthCheckServiceUnavailableCode is the HTTP code returned for type HealthCheckServiceUnavailable
const HealthCheckServiceUnavailableCode int = 503

/*HealthCheckServiceUnavailable The service is not working correctly

swagger:response healthCheckServiceUnavailable
*/
type HealthCheckServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.HealthCheckResponse `json:"body,omitempty"`
}

// NewHealthCheckServiceUnavailable creates HealthCheckServiceUnavailable with default headers values
func NewHealthCheckServiceUnavailable() *HealthCheckServiceUnavailable {
	return &HealthCheckServiceUnavailable{}
}

// WithPayload adds the payload to the health check service unavailable response
func (o *HealthCheckServiceUnavailable) WithPayload(payload *models.HealthCheckResponse) *HealthCheckServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the health check service unavailable response
func (o *HealthCheckServiceUnavailable) SetPayload(payload *models.HealthCheckResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HealthCheckServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
