// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	models "github.com/Maksim646/Tokens/internal/api/definition"
)

// NewPostAuthRefreshParams creates a new PostAuthRefreshParams object
//
// There are no default values defined in the spec.
func NewPostAuthRefreshParams() PostAuthRefreshParams {

	return PostAuthRefreshParams{}
}

// PostAuthRefreshParams contains all the bound params for the post auth refresh operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostAuthRefresh
type PostAuthRefreshParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Refresh Token Body
	  Required: true
	  In: body
	*/
	RefreshToken *models.RefreshTokenBody
	/*Client's IP address
	  In: header
	*/
	XRealIP *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostAuthRefreshParams() beforehand.
func (o *PostAuthRefreshParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.RefreshTokenBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("refreshToken", "body", ""))
			} else {
				res = append(res, errors.NewParseError("refreshToken", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.RefreshToken = &body
			}
		}
	} else {
		res = append(res, errors.Required("refreshToken", "body", ""))
	}

	if err := o.bindXRealIP(r.Header[http.CanonicalHeaderKey("X-Real-IP")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRealIP binds and validates parameter XRealIP from header.
func (o *PostAuthRefreshParams) bindXRealIP(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.XRealIP = &raw

	return nil
}
