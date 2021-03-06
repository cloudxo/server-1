// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/freecloudio/server/models"
)

// SearchFileHandlerFunc turns a function with the right signature into a search file handler
type SearchFileHandlerFunc func(SearchFileParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchFileHandlerFunc) Handle(params SearchFileParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// SearchFileHandler interface for that can handle valid search file params
type SearchFileHandler interface {
	Handle(SearchFileParams, *models.Principal) middleware.Responder
}

// NewSearchFile creates a new http.Handler for the search file operation
func NewSearchFile(ctx *middleware.Context, handler SearchFileHandler) *SearchFile {
	return &SearchFile{Context: ctx, Handler: handler}
}

/*SearchFile swagger:route GET /file/search file searchFile

Search files/folders

*/
type SearchFile struct {
	Context *middleware.Context
	Handler SearchFileHandler
}

func (o *SearchFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSearchFileParams()

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
