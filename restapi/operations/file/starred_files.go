// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// StarredFilesHandlerFunc turns a function with the right signature into a starred files handler
type StarredFilesHandlerFunc func(StarredFilesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn StarredFilesHandlerFunc) Handle(params StarredFilesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// StarredFilesHandler interface for that can handle valid starred files params
type StarredFilesHandler interface {
	Handle(StarredFilesParams, interface{}) middleware.Responder
}

// NewStarredFiles creates a new http.Handler for the starred files operation
func NewStarredFiles(ctx *middleware.Context, handler StarredFilesHandler) *StarredFiles {
	return &StarredFiles{Context: ctx, Handler: handler}
}

/*StarredFiles swagger:route GET /starred file starredFiles

Get starred files/folders

*/
type StarredFiles struct {
	Context *middleware.Context
	Handler StarredFilesHandler
}

func (o *StarredFiles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewStarredFilesParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
