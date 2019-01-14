// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/freecloudio/server/models"
)

// UpdateCurrentUserHandlerFunc turns a function with the right signature into a update current user handler
type UpdateCurrentUserHandlerFunc func(UpdateCurrentUserParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateCurrentUserHandlerFunc) Handle(params UpdateCurrentUserParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UpdateCurrentUserHandler interface for that can handle valid update current user params
type UpdateCurrentUserHandler interface {
	Handle(UpdateCurrentUserParams, *models.Principal) middleware.Responder
}

// NewUpdateCurrentUser creates a new http.Handler for the update current user operation
func NewUpdateCurrentUser(ctx *middleware.Context, handler UpdateCurrentUserHandler) *UpdateCurrentUser {
	return &UpdateCurrentUser{Context: ctx, Handler: handler}
}

/*UpdateCurrentUser swagger:route PATCH /user/me user updateCurrentUser

Modify current user

*/
type UpdateCurrentUser struct {
	Context *middleware.Context
	Handler UpdateCurrentUserHandler
}

func (o *UpdateCurrentUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateCurrentUserParams()

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
