// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// AddNamespaceReader is a Reader for the AddNamespace structure.
type AddNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddNamespaceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddNamespaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddNamespaceCreated creates a AddNamespaceCreated with default headers values
func NewAddNamespaceCreated() *AddNamespaceCreated {
	return &AddNamespaceCreated{}
}

/*AddNamespaceCreated handles this case with default header values.

Returns the created namespace
*/
type AddNamespaceCreated struct {
	Payload *inventory_models.Namespace
}

func (o *AddNamespaceCreated) Error() string {
	return fmt.Sprintf("[POST /v1/namespaces][%d] addNamespaceCreated  %+v", 201, o.Payload)
}

func (o *AddNamespaceCreated) GetPayload() *inventory_models.Namespace {
	return o.Payload
}

func (o *AddNamespaceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.Namespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceBadRequest creates a AddNamespaceBadRequest with default headers values
func NewAddNamespaceBadRequest() *AddNamespaceBadRequest {
	return &AddNamespaceBadRequest{}
}

/*AddNamespaceBadRequest handles this case with default header values.

If the namespace is invalid
*/
type AddNamespaceBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/namespaces][%d] addNamespaceBadRequest  %+v", 400, o.Payload)
}

func (o *AddNamespaceBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddNamespaceDefault creates a AddNamespaceDefault with default headers values
func NewAddNamespaceDefault(code int) *AddNamespaceDefault {
	return &AddNamespaceDefault{
		_statusCode: code,
	}
}

/*AddNamespaceDefault handles this case with default header values.

If there is an error processing the namespace
*/
type AddNamespaceDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add namespace default response
func (o *AddNamespaceDefault) Code() int {
	return o._statusCode
}

func (o *AddNamespaceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/namespaces][%d] addNamespace default  %+v", o._statusCode, o.Payload)
}

func (o *AddNamespaceDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddNamespaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
