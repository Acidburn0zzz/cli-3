// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// AddFormatReader is a Reader for the AddFormat structure.
type AddFormatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddFormatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddFormatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddFormatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddFormatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddFormatConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddFormatInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddFormatOK creates a AddFormatOK with default headers values
func NewAddFormatOK() *AddFormatOK {
	return &AddFormatOK{}
}

/* AddFormatOK describes a response with status code 200, with default header values.

Distro Added
*/
type AddFormatOK struct {
	Payload *mono_models.Format
}

func (o *AddFormatOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats][%d] addFormatOK  %+v", 200, o.Payload)
}
func (o *AddFormatOK) GetPayload() *mono_models.Format {
	return o.Payload
}

func (o *AddFormatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Format)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddFormatBadRequest creates a AddFormatBadRequest with default headers values
func NewAddFormatBadRequest() *AddFormatBadRequest {
	return &AddFormatBadRequest{}
}

/* AddFormatBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddFormatBadRequest struct {
	Payload *mono_models.Message
}

func (o *AddFormatBadRequest) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats][%d] addFormatBadRequest  %+v", 400, o.Payload)
}
func (o *AddFormatBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddFormatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddFormatForbidden creates a AddFormatForbidden with default headers values
func NewAddFormatForbidden() *AddFormatForbidden {
	return &AddFormatForbidden{}
}

/* AddFormatForbidden describes a response with status code 403, with default header values.

Unauthorized
*/
type AddFormatForbidden struct {
	Payload *mono_models.Message
}

func (o *AddFormatForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats][%d] addFormatForbidden  %+v", 403, o.Payload)
}
func (o *AddFormatForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddFormatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddFormatConflict creates a AddFormatConflict with default headers values
func NewAddFormatConflict() *AddFormatConflict {
	return &AddFormatConflict{}
}

/* AddFormatConflict describes a response with status code 409, with default header values.

Conflict
*/
type AddFormatConflict struct {
	Payload *mono_models.Message
}

func (o *AddFormatConflict) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats][%d] addFormatConflict  %+v", 409, o.Payload)
}
func (o *AddFormatConflict) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddFormatConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddFormatInternalServerError creates a AddFormatInternalServerError with default headers values
func NewAddFormatInternalServerError() *AddFormatInternalServerError {
	return &AddFormatInternalServerError{}
}

/* AddFormatInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AddFormatInternalServerError struct {
	Payload *mono_models.Message
}

func (o *AddFormatInternalServerError) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats][%d] addFormatInternalServerError  %+v", 500, o.Payload)
}
func (o *AddFormatInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddFormatInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
