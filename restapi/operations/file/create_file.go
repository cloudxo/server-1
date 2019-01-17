// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/freecloudio/server/models"
)

// CreateFileHandlerFunc turns a function with the right signature into a create file handler
type CreateFileHandlerFunc func(CreateFileParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateFileHandlerFunc) Handle(params CreateFileParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// CreateFileHandler interface for that can handle valid create file params
type CreateFileHandler interface {
	Handle(CreateFileParams, *models.Principal) middleware.Responder
}

// NewCreateFile creates a new http.Handler for the create file operation
func NewCreateFile(ctx *middleware.Context, handler CreateFileHandler) *CreateFile {
	return &CreateFile{Context: ctx, Handler: handler}
}

/*CreateFile swagger:route POST /file file createFile

Create new file/folder

*/
type CreateFile struct {
	Context *middleware.Context
	Handler CreateFileHandler
}

func (o *CreateFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateFileParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}