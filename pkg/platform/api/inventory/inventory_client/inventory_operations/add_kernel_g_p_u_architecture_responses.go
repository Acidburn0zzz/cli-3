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

// AddKernelGPUArchitectureReader is a Reader for the AddKernelGPUArchitecture structure.
type AddKernelGPUArchitectureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddKernelGPUArchitectureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddKernelGPUArchitectureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddKernelGPUArchitectureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddKernelGPUArchitectureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddKernelGPUArchitectureOK creates a AddKernelGPUArchitectureOK with default headers values
func NewAddKernelGPUArchitectureOK() *AddKernelGPUArchitectureOK {
	return &AddKernelGPUArchitectureOK{}
}

/*AddKernelGPUArchitectureOK handles this case with default header values.

The GPU architecture added to the kernel
*/
type AddKernelGPUArchitectureOK struct {
	Payload *inventory_models.GpuArchitecture
}

func (o *AddKernelGPUArchitectureOK) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/gpu-architectures][%d] addKernelGPUArchitectureOK  %+v", 200, o.Payload)
}

func (o *AddKernelGPUArchitectureOK) GetPayload() *inventory_models.GpuArchitecture {
	return o.Payload
}

func (o *AddKernelGPUArchitectureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.GpuArchitecture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKernelGPUArchitectureBadRequest creates a AddKernelGPUArchitectureBadRequest with default headers values
func NewAddKernelGPUArchitectureBadRequest() *AddKernelGPUArchitectureBadRequest {
	return &AddKernelGPUArchitectureBadRequest{}
}

/*AddKernelGPUArchitectureBadRequest handles this case with default header values.

If the GPU architecture ID doesn't exist
*/
type AddKernelGPUArchitectureBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddKernelGPUArchitectureBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/gpu-architectures][%d] addKernelGPUArchitectureBadRequest  %+v", 400, o.Payload)
}

func (o *AddKernelGPUArchitectureBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddKernelGPUArchitectureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKernelGPUArchitectureDefault creates a AddKernelGPUArchitectureDefault with default headers values
func NewAddKernelGPUArchitectureDefault(code int) *AddKernelGPUArchitectureDefault {
	return &AddKernelGPUArchitectureDefault{
		_statusCode: code,
	}
}

/*AddKernelGPUArchitectureDefault handles this case with default header values.

generic error response
*/
type AddKernelGPUArchitectureDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add kernel g p u architecture default response
func (o *AddKernelGPUArchitectureDefault) Code() int {
	return o._statusCode
}

func (o *AddKernelGPUArchitectureDefault) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/gpu-architectures][%d] addKernelGPUArchitecture default  %+v", o._statusCode, o.Payload)
}

func (o *AddKernelGPUArchitectureDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddKernelGPUArchitectureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
