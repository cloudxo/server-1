// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/freecloudio/freecloud/models"
)

// UpdateUserByIDOKCode is the HTTP code returned for type UpdateUserByIDOK
const UpdateUserByIDOKCode int = 200

/*UpdateUserByIDOK Success

swagger:response updateUserByIdOK
*/
type UpdateUserByIDOK struct {
}

// NewUpdateUserByIDOK creates UpdateUserByIDOK with default headers values
func NewUpdateUserByIDOK() *UpdateUserByIDOK {

	return &UpdateUserByIDOK{}
}

// WriteResponse to the client
func (o *UpdateUserByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*UpdateUserByIDDefault Unexpected error

swagger:response updateUserByIdDefault
*/
type UpdateUserByIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserByIDDefault creates UpdateUserByIDDefault with default headers values
func NewUpdateUserByIDDefault(code int) *UpdateUserByIDDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateUserByIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update user by ID default response
func (o *UpdateUserByIDDefault) WithStatusCode(code int) *UpdateUserByIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update user by ID default response
func (o *UpdateUserByIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update user by ID default response
func (o *UpdateUserByIDDefault) WithPayload(payload *models.Error) *UpdateUserByIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user by ID default response
func (o *UpdateUserByIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserByIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
