// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/freecloudio/freecloud/models"
)

// SearchFileOKCode is the HTTP code returned for type SearchFileOK
const SearchFileOKCode int = 200

/*SearchFileOK Search results

swagger:response searchFileOK
*/
type SearchFileOK struct {

	/*
	  In: Body
	*/
	Payload *models.FileList `json:"body,omitempty"`
}

// NewSearchFileOK creates SearchFileOK with default headers values
func NewSearchFileOK() *SearchFileOK {

	return &SearchFileOK{}
}

// WithPayload adds the payload to the search file o k response
func (o *SearchFileOK) WithPayload(payload *models.FileList) *SearchFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search file o k response
func (o *SearchFileOK) SetPayload(payload *models.FileList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*SearchFileDefault Unexpected error

swagger:response searchFileDefault
*/
type SearchFileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSearchFileDefault creates SearchFileDefault with default headers values
func NewSearchFileDefault(code int) *SearchFileDefault {
	if code <= 0 {
		code = 500
	}

	return &SearchFileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the search file default response
func (o *SearchFileDefault) WithStatusCode(code int) *SearchFileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the search file default response
func (o *SearchFileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the search file default response
func (o *SearchFileDefault) WithPayload(payload *models.Error) *SearchFileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search file default response
func (o *SearchFileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchFileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
