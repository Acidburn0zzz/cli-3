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

// AddGPUArchitectureRevisionReader is a Reader for the AddGPUArchitectureRevision structure.
type AddGPUArchitectureRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddGPUArchitectureRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddGPUArchitectureRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddGPUArchitectureRevisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddGPUArchitectureRevisionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddGPUArchitectureRevisionOK creates a AddGPUArchitectureRevisionOK with default headers values
func NewAddGPUArchitectureRevisionOK() *AddGPUArchitectureRevisionOK {
	return &AddGPUArchitectureRevisionOK{}
}

/*AddGPUArchitectureRevisionOK handles this case with default header values.

The updated state of the GPU architecture
*/
type AddGPUArchitectureRevisionOK struct {
	Payload *inventory_models.GpuArchitecture
}

func (o *AddGPUArchitectureRevisionOK) Error() string {
	return fmt.Sprintf("[POST /v1/gpu-architectures/{gpu_architecture_id}/revisions][%d] addGPUArchitectureRevisionOK  %+v", 200, o.Payload)
}

func (o *AddGPUArchitectureRevisionOK) GetPayload() *inventory_models.GpuArchitecture {
	return o.Payload
}

func (o *AddGPUArchitectureRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.GpuArchitecture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGPUArchitectureRevisionBadRequest creates a AddGPUArchitectureRevisionBadRequest with default headers values
func NewAddGPUArchitectureRevisionBadRequest() *AddGPUArchitectureRevisionBadRequest {
	return &AddGPUArchitectureRevisionBadRequest{}
}

/*AddGPUArchitectureRevisionBadRequest handles this case with default header values.

If the GPU architecture revision is invalid
*/
type AddGPUArchitectureRevisionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddGPUArchitectureRevisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/gpu-architectures/{gpu_architecture_id}/revisions][%d] addGPUArchitectureRevisionBadRequest  %+v", 400, o.Payload)
}

func (o *AddGPUArchitectureRevisionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddGPUArchitectureRevisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddGPUArchitectureRevisionDefault creates a AddGPUArchitectureRevisionDefault with default headers values
func NewAddGPUArchitectureRevisionDefault(code int) *AddGPUArchitectureRevisionDefault {
	return &AddGPUArchitectureRevisionDefault{
		_statusCode: code,
	}
}

/*AddGPUArchitectureRevisionDefault handles this case with default header values.

If there is an error processing the request
*/
type AddGPUArchitectureRevisionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add g p u architecture revision default response
func (o *AddGPUArchitectureRevisionDefault) Code() int {
	return o._statusCode
}

func (o *AddGPUArchitectureRevisionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/gpu-architectures/{gpu_architecture_id}/revisions][%d] addGPUArchitectureRevision default  %+v", o._statusCode, o.Payload)
}

func (o *AddGPUArchitectureRevisionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddGPUArchitectureRevisionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
