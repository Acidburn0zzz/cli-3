// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetLoginJwtTokenReader is a Reader for the GetLoginJwtToken structure.
type GetLoginJwtTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoginJwtTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetLoginJwtTokenFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetLoginJwtTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLoginJwtTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLoginJwtTokenFound creates a GetLoginJwtTokenFound with default headers values
func NewGetLoginJwtTokenFound() *GetLoginJwtTokenFound {
	return &GetLoginJwtTokenFound{}
}

/* GetLoginJwtTokenFound describes a response with status code 302, with default header values.

Found
*/
type GetLoginJwtTokenFound struct {
}

func (o *GetLoginJwtTokenFound) Error() string {
	return fmt.Sprintf("[GET /login/jwt/{token}][%d] getLoginJwtTokenFound ", 302)
}

func (o *GetLoginJwtTokenFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLoginJwtTokenBadRequest creates a GetLoginJwtTokenBadRequest with default headers values
func NewGetLoginJwtTokenBadRequest() *GetLoginJwtTokenBadRequest {
	return &GetLoginJwtTokenBadRequest{}
}

/* GetLoginJwtTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLoginJwtTokenBadRequest struct {
	Payload *mono_models.Message
}

func (o *GetLoginJwtTokenBadRequest) Error() string {
	return fmt.Sprintf("[GET /login/jwt/{token}][%d] getLoginJwtTokenBadRequest  %+v", 400, o.Payload)
}
func (o *GetLoginJwtTokenBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetLoginJwtTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoginJwtTokenInternalServerError creates a GetLoginJwtTokenInternalServerError with default headers values
func NewGetLoginJwtTokenInternalServerError() *GetLoginJwtTokenInternalServerError {
	return &GetLoginJwtTokenInternalServerError{}
}

/* GetLoginJwtTokenInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetLoginJwtTokenInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetLoginJwtTokenInternalServerError) Error() string {
	return fmt.Sprintf("[GET /login/jwt/{token}][%d] getLoginJwtTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLoginJwtTokenInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetLoginJwtTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
