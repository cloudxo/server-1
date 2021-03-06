// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/freecloudio/server/models"
)

// UpdateCurrentUserOKCode is the HTTP code returned for type UpdateCurrentUserOK
const UpdateCurrentUserOKCode int = 200

/*UpdateCurrentUserOK Success

swagger:response updateCurrentUserOK
*/
type UpdateCurrentUserOK struct {
}

// NewUpdateCurrentUserOK creates UpdateCurrentUserOK with default headers values
func NewUpdateCurrentUserOK() *UpdateCurrentUserOK {

	return &UpdateCurrentUserOK{}
}

// WriteResponse to the client
func (o *UpdateCurrentUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*UpdateCurrentUserDefault Unexpected error

swagger:response updateCurrentUserDefault
*/
type UpdateCurrentUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateCurrentUserDefault creates UpdateCurrentUserDefault with default headers values
func NewUpdateCurrentUserDefault(code int) *UpdateCurrentUserDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateCurrentUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update current user default response
func (o *UpdateCurrentUserDefault) WithStatusCode(code int) *UpdateCurrentUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update current user default response
func (o *UpdateCurrentUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update current user default response
func (o *UpdateCurrentUserDefault) WithPayload(payload *models.Error) *UpdateCurrentUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update current user default response
func (o *UpdateCurrentUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateCurrentUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
