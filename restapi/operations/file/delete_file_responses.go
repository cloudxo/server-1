// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/freecloudio/freecloud/models"
)

// DeleteFileOKCode is the HTTP code returned for type DeleteFileOK
const DeleteFileOKCode int = 200

/*DeleteFileOK Success

swagger:response deleteFileOK
*/
type DeleteFileOK struct {
}

// NewDeleteFileOK creates DeleteFileOK with default headers values
func NewDeleteFileOK() *DeleteFileOK {

	return &DeleteFileOK{}
}

// WriteResponse to the client
func (o *DeleteFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeleteFileDefault Unexpected error

swagger:response deleteFileDefault
*/
type DeleteFileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteFileDefault creates DeleteFileDefault with default headers values
func NewDeleteFileDefault(code int) *DeleteFileDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteFileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete file default response
func (o *DeleteFileDefault) WithStatusCode(code int) *DeleteFileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete file default response
func (o *DeleteFileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete file default response
func (o *DeleteFileDefault) WithPayload(payload *models.Error) *DeleteFileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete file default response
func (o *DeleteFileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
