// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetNextMutationIDReader is a Reader for the GetNextMutationID structure.
type GetNextMutationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNextMutationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNextMutationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetNextMutationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetNextMutationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetNextMutationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNextMutationIDOK creates a GetNextMutationIDOK with default headers values
func NewGetNextMutationIDOK() *GetNextMutationIDOK {
	return &GetNextMutationIDOK{}
}

/*GetNextMutationIDOK handles this case with default header values.

Success
*/
type GetNextMutationIDOK struct {
	Payload string
}

func (o *GetNextMutationIDOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationIdentifier}/nextMutationID][%d] getNextMutationIdOK  %+v", 200, o.Payload)
}

func (o *GetNextMutationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNextMutationIDForbidden creates a GetNextMutationIDForbidden with default headers values
func NewGetNextMutationIDForbidden() *GetNextMutationIDForbidden {
	return &GetNextMutationIDForbidden{}
}

/*GetNextMutationIDForbidden handles this case with default header values.

Unauthorized
*/
type GetNextMutationIDForbidden struct {
	Payload *mono_models.Message
}

func (o *GetNextMutationIDForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationIdentifier}/nextMutationID][%d] getNextMutationIdForbidden  %+v", 403, o.Payload)
}

func (o *GetNextMutationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNextMutationIDNotFound creates a GetNextMutationIDNotFound with default headers values
func NewGetNextMutationIDNotFound() *GetNextMutationIDNotFound {
	return &GetNextMutationIDNotFound{}
}

/*GetNextMutationIDNotFound handles this case with default header values.

Not Found
*/
type GetNextMutationIDNotFound struct {
	Payload *mono_models.Message
}

func (o *GetNextMutationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationIdentifier}/nextMutationID][%d] getNextMutationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetNextMutationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNextMutationIDInternalServerError creates a GetNextMutationIDInternalServerError with default headers values
func NewGetNextMutationIDInternalServerError() *GetNextMutationIDInternalServerError {
	return &GetNextMutationIDInternalServerError{}
}

/*GetNextMutationIDInternalServerError handles this case with default header values.

Server Error
*/
type GetNextMutationIDInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetNextMutationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationIdentifier}/nextMutationID][%d] getNextMutationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNextMutationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}