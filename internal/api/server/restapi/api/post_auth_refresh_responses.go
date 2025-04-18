// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Maksim646/tokens/internal/api/definition"
)

// PostAuthRefreshOKCode is the HTTP code returned for type PostAuthRefreshOK
const PostAuthRefreshOKCode int = 200

/*
PostAuthRefreshOK Successful Tokens Response

swagger:response postAuthRefreshOK
*/
type PostAuthRefreshOK struct {

	/*
	  In: Body
	*/
	Payload *models.AccessTokenBody `json:"body,omitempty"`
}

// NewPostAuthRefreshOK creates PostAuthRefreshOK with default headers values
func NewPostAuthRefreshOK() *PostAuthRefreshOK {

	return &PostAuthRefreshOK{}
}

// WithPayload adds the payload to the post auth refresh o k response
func (o *PostAuthRefreshOK) WithPayload(payload *models.AccessTokenBody) *PostAuthRefreshOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth refresh o k response
func (o *PostAuthRefreshOK) SetPayload(payload *models.AccessTokenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthRefreshOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthRefreshBadRequestCode is the HTTP code returned for type PostAuthRefreshBadRequest
const PostAuthRefreshBadRequestCode int = 400

/*
PostAuthRefreshBadRequest Bad request

swagger:response postAuthRefreshBadRequest
*/
type PostAuthRefreshBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostAuthRefreshBadRequest creates PostAuthRefreshBadRequest with default headers values
func NewPostAuthRefreshBadRequest() *PostAuthRefreshBadRequest {

	return &PostAuthRefreshBadRequest{}
}

// WithPayload adds the payload to the post auth refresh bad request response
func (o *PostAuthRefreshBadRequest) WithPayload(payload *models.Error) *PostAuthRefreshBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth refresh bad request response
func (o *PostAuthRefreshBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthRefreshBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthRefreshUnauthorizedCode is the HTTP code returned for type PostAuthRefreshUnauthorized
const PostAuthRefreshUnauthorizedCode int = 401

/*
PostAuthRefreshUnauthorized Unauthorized

swagger:response postAuthRefreshUnauthorized
*/
type PostAuthRefreshUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostAuthRefreshUnauthorized creates PostAuthRefreshUnauthorized with default headers values
func NewPostAuthRefreshUnauthorized() *PostAuthRefreshUnauthorized {

	return &PostAuthRefreshUnauthorized{}
}

// WithPayload adds the payload to the post auth refresh unauthorized response
func (o *PostAuthRefreshUnauthorized) WithPayload(payload *models.Error) *PostAuthRefreshUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth refresh unauthorized response
func (o *PostAuthRefreshUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthRefreshUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthRefreshForbiddenCode is the HTTP code returned for type PostAuthRefreshForbidden
const PostAuthRefreshForbiddenCode int = 403

/*
PostAuthRefreshForbidden Invalid IP or token mismatch

swagger:response postAuthRefreshForbidden
*/
type PostAuthRefreshForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostAuthRefreshForbidden creates PostAuthRefreshForbidden with default headers values
func NewPostAuthRefreshForbidden() *PostAuthRefreshForbidden {

	return &PostAuthRefreshForbidden{}
}

// WithPayload adds the payload to the post auth refresh forbidden response
func (o *PostAuthRefreshForbidden) WithPayload(payload *models.Error) *PostAuthRefreshForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth refresh forbidden response
func (o *PostAuthRefreshForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthRefreshForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthRefreshConflictCode is the HTTP code returned for type PostAuthRefreshConflict
const PostAuthRefreshConflictCode int = 409

/*
PostAuthRefreshConflict Refresh token reuse attempt

swagger:response postAuthRefreshConflict
*/
type PostAuthRefreshConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostAuthRefreshConflict creates PostAuthRefreshConflict with default headers values
func NewPostAuthRefreshConflict() *PostAuthRefreshConflict {

	return &PostAuthRefreshConflict{}
}

// WithPayload adds the payload to the post auth refresh conflict response
func (o *PostAuthRefreshConflict) WithPayload(payload *models.Error) *PostAuthRefreshConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth refresh conflict response
func (o *PostAuthRefreshConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthRefreshConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAuthRefreshInternalServerErrorCode is the HTTP code returned for type PostAuthRefreshInternalServerError
const PostAuthRefreshInternalServerErrorCode int = 500

/*
PostAuthRefreshInternalServerError Internal server error

swagger:response postAuthRefreshInternalServerError
*/
type PostAuthRefreshInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostAuthRefreshInternalServerError creates PostAuthRefreshInternalServerError with default headers values
func NewPostAuthRefreshInternalServerError() *PostAuthRefreshInternalServerError {

	return &PostAuthRefreshInternalServerError{}
}

// WithPayload adds the payload to the post auth refresh internal server error response
func (o *PostAuthRefreshInternalServerError) WithPayload(payload *models.Error) *PostAuthRefreshInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth refresh internal server error response
func (o *PostAuthRefreshInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthRefreshInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
