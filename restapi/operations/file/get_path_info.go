// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/freecloudio/freecloud/models"
)

// GetPathInfoHandlerFunc turns a function with the right signature into a get path info handler
type GetPathInfoHandlerFunc func(GetPathInfoParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPathInfoHandlerFunc) Handle(params GetPathInfoParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetPathInfoHandler interface for that can handle valid get path info params
type GetPathInfoHandler interface {
	Handle(GetPathInfoParams, *models.Principal) middleware.Responder
}

// NewGetPathInfo creates a new http.Handler for the get path info operation
func NewGetPathInfo(ctx *middleware.Context, handler GetPathInfoHandler) *GetPathInfo {
	return &GetPathInfo{Context: ctx, Handler: handler}
}

/*GetPathInfo swagger:route GET /file file getPathInfo

Get pathInfo of requested path

*/
type GetPathInfo struct {
	Context *middleware.Context
	Handler GetPathInfoHandler
}

func (o *GetPathInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPathInfoParams()

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
