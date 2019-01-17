// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/freecloudio/server/models"
)

// ZipFilesOKCode is the HTTP code returned for type ZipFilesOK
const ZipFilesOKCode int = 200

/*ZipFilesOK Success

swagger:response zipFilesOK
*/
type ZipFilesOK struct {

	/*
	  In: Body
	*/
	Payload *models.Path `json:"body,omitempty"`
}

// NewZipFilesOK creates ZipFilesOK with default headers values
func NewZipFilesOK() *ZipFilesOK {

	return &ZipFilesOK{}
}

// WithPayload adds the payload to the zip files o k response
func (o *ZipFilesOK) WithPayload(payload *models.Path) *ZipFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zip files o k response
func (o *ZipFilesOK) SetPayload(payload *models.Path) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ZipFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ZipFilesDefault Unexpected error

swagger:response zipFilesDefault
*/
type ZipFilesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewZipFilesDefault creates ZipFilesDefault with default headers values
func NewZipFilesDefault(code int) *ZipFilesDefault {
	if code <= 0 {
		code = 500
	}

	return &ZipFilesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the zip files default response
func (o *ZipFilesDefault) WithStatusCode(code int) *ZipFilesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the zip files default response
func (o *ZipFilesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the zip files default response
func (o *ZipFilesDefault) WithPayload(payload *models.Error) *ZipFilesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the zip files default response
func (o *ZipFilesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ZipFilesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}