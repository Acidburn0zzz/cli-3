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

// GetProjectByIDReader is a Reader for the GetProjectByID structure.
type GetProjectByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProjectByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectByIDOK creates a GetProjectByIDOK with default headers values
func NewGetProjectByIDOK() *GetProjectByIDOK {
	return &GetProjectByIDOK{}
}

/* GetProjectByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetProjectByIDOK struct {
	Payload *mono_models.Project
}

func (o *GetProjectByIDOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectID}][%d] getProjectByIdOK  %+v", 200, o.Payload)
}
func (o *GetProjectByIDOK) GetPayload() *mono_models.Project {
	return o.Payload
}

func (o *GetProjectByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectByIDNotFound creates a GetProjectByIDNotFound with default headers values
func NewGetProjectByIDNotFound() *GetProjectByIDNotFound {
	return &GetProjectByIDNotFound{}
}

/* GetProjectByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetProjectByIDNotFound struct {
	Payload *mono_models.Message
}

func (o *GetProjectByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{projectID}][%d] getProjectByIdNotFound  %+v", 404, o.Payload)
}
func (o *GetProjectByIDNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetProjectByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectByIDInternalServerError creates a GetProjectByIDInternalServerError with default headers values
func NewGetProjectByIDInternalServerError() *GetProjectByIDInternalServerError {
	return &GetProjectByIDInternalServerError{}
}

/* GetProjectByIDInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetProjectByIDInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetProjectByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{projectID}][%d] getProjectByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProjectByIDInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetProjectByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
