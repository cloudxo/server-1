// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/freecloudio/server/models"
)

// NewLoginParams creates a new LoginParams object
// no default values defined in spec.
func NewLoginParams() LoginParams {

	return LoginParams{}
}

// LoginParams contains all the bound params for the login operation
// typically these are obtained from a http.Request
//
// swagger:parameters login
type LoginParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Credentials for login
	  Required: true
	  In: body
	*/
	Credentials *models.LoginData
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewLoginParams() beforehand.
func (o *LoginParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.LoginData
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("credentials", "body"))
			} else {
				res = append(res, errors.NewParseError("credentials", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Credentials = &body
			}
		}
	} else {
		res = append(res, errors.Required("credentials", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
