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

// AddCPUArchitectureRevisionReader is a Reader for the AddCPUArchitectureRevision structure.
type AddCPUArchitectureRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddCPUArchitectureRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddCPUArchitectureRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddCPUArchitectureRevisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddCPUArchitectureRevisionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddCPUArchitectureRevisionOK creates a AddCPUArchitectureRevisionOK with default headers values
func NewAddCPUArchitectureRevisionOK() *AddCPUArchitectureRevisionOK {
	return &AddCPUArchitectureRevisionOK{}
}

/* AddCPUArchitectureRevisionOK describes a response with status code 200, with default header values.

The updated state of the CPU architecture
*/
type AddCPUArchitectureRevisionOK struct {
	Payload *inventory_models.CPUArchitecture
}

func (o *AddCPUArchitectureRevisionOK) Error() string {
	return fmt.Sprintf("[POST /v1/cpu-architectures/{cpu_architecture_id}/revisions][%d] addCpuArchitectureRevisionOK  %+v", 200, o.Payload)
}
func (o *AddCPUArchitectureRevisionOK) GetPayload() *inventory_models.CPUArchitecture {
	return o.Payload
}

func (o *AddCPUArchitectureRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.CPUArchitecture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCPUArchitectureRevisionBadRequest creates a AddCPUArchitectureRevisionBadRequest with default headers values
func NewAddCPUArchitectureRevisionBadRequest() *AddCPUArchitectureRevisionBadRequest {
	return &AddCPUArchitectureRevisionBadRequest{}
}

/* AddCPUArchitectureRevisionBadRequest describes a response with status code 400, with default header values.

If the CPU architecture revision is invalid
*/
type AddCPUArchitectureRevisionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddCPUArchitectureRevisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/cpu-architectures/{cpu_architecture_id}/revisions][%d] addCpuArchitectureRevisionBadRequest  %+v", 400, o.Payload)
}
func (o *AddCPUArchitectureRevisionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddCPUArchitectureRevisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddCPUArchitectureRevisionDefault creates a AddCPUArchitectureRevisionDefault with default headers values
func NewAddCPUArchitectureRevisionDefault(code int) *AddCPUArchitectureRevisionDefault {
	return &AddCPUArchitectureRevisionDefault{
		_statusCode: code,
	}
}

/* AddCPUArchitectureRevisionDefault describes a response with status code -1, with default header values.

If there is an error processing the request
*/
type AddCPUArchitectureRevisionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add Cpu architecture revision default response
func (o *AddCPUArchitectureRevisionDefault) Code() int {
	return o._statusCode
}

func (o *AddCPUArchitectureRevisionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/cpu-architectures/{cpu_architecture_id}/revisions][%d] addCpuArchitectureRevision default  %+v", o._statusCode, o.Payload)
}
func (o *AddCPUArchitectureRevisionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddCPUArchitectureRevisionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
