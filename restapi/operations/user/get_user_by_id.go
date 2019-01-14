// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/freecloudio/freecloud/models"
)

// GetUserByIDHandlerFunc turns a function with the right signature into a get user by ID handler
type GetUserByIDHandlerFunc func(GetUserByIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserByIDHandlerFunc) Handle(params GetUserByIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetUserByIDHandler interface for that can handle valid get user by ID params
type GetUserByIDHandler interface {
	Handle(GetUserByIDParams, *models.Principal) middleware.Responder
}

// NewGetUserByID creates a new http.Handler for the get user by ID operation
func NewGetUserByID(ctx *middleware.Context, handler GetUserByIDHandler) *GetUserByID {
	return &GetUserByID{Context: ctx, Handler: handler}
}

/*GetUserByID swagger:route GET /user/{id} user getUserById

Get specific user by id

*/
type GetUserByID struct {
	Context *middleware.Context
	Handler GetUserByIDHandler
}

func (o *GetUserByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserByIDParams()

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
