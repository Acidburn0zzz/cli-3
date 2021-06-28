// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetEmailsByUserReader is a Reader for the GetEmailsByUser structure.
type GetEmailsByUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailsByUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmailsByUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetEmailsByUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEmailsByUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEmailsByUserOK creates a GetEmailsByUserOK with default headers values
func NewGetEmailsByUserOK() *GetEmailsByUserOK {
	return &GetEmailsByUserOK{}
}

/* GetEmailsByUserOK describes a response with status code 200, with default header values.

Email records
*/
type GetEmailsByUserOK struct {
	Payload []*mono_models.Email
}

func (o *GetEmailsByUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{username}/emails][%d] getEmailsByUserOK  %+v", 200, o.Payload)
}
func (o *GetEmailsByUserOK) GetPayload() []*mono_models.Email {
	return o.Payload
}

func (o *GetEmailsByUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailsByUserForbidden creates a GetEmailsByUserForbidden with default headers values
func NewGetEmailsByUserForbidden() *GetEmailsByUserForbidden {
	return &GetEmailsByUserForbidden{}
}

/* GetEmailsByUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEmailsByUserForbidden struct {
	Payload *mono_models.Message
}

func (o *GetEmailsByUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{username}/emails][%d] getEmailsByUserForbidden  %+v", 403, o.Payload)
}
func (o *GetEmailsByUserForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetEmailsByUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailsByUserInternalServerError creates a GetEmailsByUserInternalServerError with default headers values
func NewGetEmailsByUserInternalServerError() *GetEmailsByUserInternalServerError {
	return &GetEmailsByUserInternalServerError{}
}

/* GetEmailsByUserInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetEmailsByUserInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetEmailsByUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{username}/emails][%d] getEmailsByUserInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEmailsByUserInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetEmailsByUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
