// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/Maksim646/tokens/internal/api/definition"
)

// GetAuthTokenReader is a Reader for the GetAuthToken structure.
type GetAuthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /auth/token] GetAuthToken", response, response.Code())
	}
}

// NewGetAuthTokenOK creates a GetAuthTokenOK with default headers values
func NewGetAuthTokenOK() *GetAuthTokenOK {
	return &GetAuthTokenOK{}
}

/*
GetAuthTokenOK describes a response with status code 200, with default header values.

Successful Tokens Response
*/
type GetAuthTokenOK struct {
	Payload *models.Tokens
}

// IsSuccess returns true when this get auth token o k response has a 2xx status code
func (o *GetAuthTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get auth token o k response has a 3xx status code
func (o *GetAuthTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get auth token o k response has a 4xx status code
func (o *GetAuthTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get auth token o k response has a 5xx status code
func (o *GetAuthTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get auth token o k response a status code equal to that given
func (o *GetAuthTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get auth token o k response
func (o *GetAuthTokenOK) Code() int {
	return 200
}

func (o *GetAuthTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/token][%d] getAuthTokenOK %s", 200, payload)
}

func (o *GetAuthTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/token][%d] getAuthTokenOK %s", 200, payload)
}

func (o *GetAuthTokenOK) GetPayload() *models.Tokens {
	return o.Payload
}

func (o *GetAuthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tokens)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthTokenBadRequest creates a GetAuthTokenBadRequest with default headers values
func NewGetAuthTokenBadRequest() *GetAuthTokenBadRequest {
	return &GetAuthTokenBadRequest{}
}

/*
GetAuthTokenBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAuthTokenBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get auth token bad request response has a 2xx status code
func (o *GetAuthTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get auth token bad request response has a 3xx status code
func (o *GetAuthTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get auth token bad request response has a 4xx status code
func (o *GetAuthTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get auth token bad request response has a 5xx status code
func (o *GetAuthTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get auth token bad request response a status code equal to that given
func (o *GetAuthTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get auth token bad request response
func (o *GetAuthTokenBadRequest) Code() int {
	return 400
}

func (o *GetAuthTokenBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/token][%d] getAuthTokenBadRequest %s", 400, payload)
}

func (o *GetAuthTokenBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/token][%d] getAuthTokenBadRequest %s", 400, payload)
}

func (o *GetAuthTokenBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuthTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthTokenInternalServerError creates a GetAuthTokenInternalServerError with default headers values
func NewGetAuthTokenInternalServerError() *GetAuthTokenInternalServerError {
	return &GetAuthTokenInternalServerError{}
}

/*
GetAuthTokenInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAuthTokenInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get auth token internal server error response has a 2xx status code
func (o *GetAuthTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get auth token internal server error response has a 3xx status code
func (o *GetAuthTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get auth token internal server error response has a 4xx status code
func (o *GetAuthTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get auth token internal server error response has a 5xx status code
func (o *GetAuthTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get auth token internal server error response a status code equal to that given
func (o *GetAuthTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get auth token internal server error response
func (o *GetAuthTokenInternalServerError) Code() int {
	return 500
}

func (o *GetAuthTokenInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/token][%d] getAuthTokenInternalServerError %s", 500, payload)
}

func (o *GetAuthTokenInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/token][%d] getAuthTokenInternalServerError %s", 500, payload)
}

func (o *GetAuthTokenInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuthTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
