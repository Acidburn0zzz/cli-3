// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// AddOrganizationReader is a Reader for the AddOrganization structure.
type AddOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddOrganizationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddOrganizationOK creates a AddOrganizationOK with default headers values
func NewAddOrganizationOK() *AddOrganizationOK {
	return &AddOrganizationOK{}
}

/* AddOrganizationOK describes a response with status code 200, with default header values.

Organization Created
*/
type AddOrganizationOK struct {
	Payload *mono_models.Organization
}

func (o *AddOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] addOrganizationOK  %+v", 200, o.Payload)
}
func (o *AddOrganizationOK) GetPayload() *mono_models.Organization {
	return o.Payload
}

func (o *AddOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationBadRequest creates a AddOrganizationBadRequest with default headers values
func NewAddOrganizationBadRequest() *AddOrganizationBadRequest {
	return &AddOrganizationBadRequest{}
}

/* AddOrganizationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddOrganizationBadRequest struct {
	Payload *mono_models.Message
}

func (o *AddOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] addOrganizationBadRequest  %+v", 400, o.Payload)
}
func (o *AddOrganizationBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationForbidden creates a AddOrganizationForbidden with default headers values
func NewAddOrganizationForbidden() *AddOrganizationForbidden {
	return &AddOrganizationForbidden{}
}

/* AddOrganizationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddOrganizationForbidden struct {
	Payload *mono_models.Message
}

func (o *AddOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] addOrganizationForbidden  %+v", 403, o.Payload)
}
func (o *AddOrganizationForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationConflict creates a AddOrganizationConflict with default headers values
func NewAddOrganizationConflict() *AddOrganizationConflict {
	return &AddOrganizationConflict{}
}

/* AddOrganizationConflict describes a response with status code 409, with default header values.

Conflict
*/
type AddOrganizationConflict struct {
	Payload *mono_models.Message
}

func (o *AddOrganizationConflict) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] addOrganizationConflict  %+v", 409, o.Payload)
}
func (o *AddOrganizationConflict) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddOrganizationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrganizationInternalServerError creates a AddOrganizationInternalServerError with default headers values
func NewAddOrganizationInternalServerError() *AddOrganizationInternalServerError {
	return &AddOrganizationInternalServerError{}
}

/* AddOrganizationInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AddOrganizationInternalServerError struct {
	Payload *mono_models.Message
}

func (o *AddOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /organizations][%d] addOrganizationInternalServerError  %+v", 500, o.Payload)
}
func (o *AddOrganizationInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
