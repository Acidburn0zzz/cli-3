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

// GetCPUArchitectureReader is a Reader for the GetCPUArchitecture structure.
type GetCPUArchitectureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCPUArchitectureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCPUArchitectureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCPUArchitectureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCPUArchitectureOK creates a GetCPUArchitectureOK with default headers values
func NewGetCPUArchitectureOK() *GetCPUArchitectureOK {
	return &GetCPUArchitectureOK{}
}

/* GetCPUArchitectureOK describes a response with status code 200, with default header values.

The retrieved CPU architecture
*/
type GetCPUArchitectureOK struct {
	Payload *inventory_models.CPUArchitecture
}

func (o *GetCPUArchitectureOK) Error() string {
	return fmt.Sprintf("[GET /v1/cpu-architectures/{cpu_architecture_id}][%d] getCpuArchitectureOK  %+v", 200, o.Payload)
}
func (o *GetCPUArchitectureOK) GetPayload() *inventory_models.CPUArchitecture {
	return o.Payload
}

func (o *GetCPUArchitectureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.CPUArchitecture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCPUArchitectureDefault creates a GetCPUArchitectureDefault with default headers values
func NewGetCPUArchitectureDefault(code int) *GetCPUArchitectureDefault {
	return &GetCPUArchitectureDefault{
		_statusCode: code,
	}
}

/* GetCPUArchitectureDefault describes a response with status code -1, with default header values.

generic error response
*/
type GetCPUArchitectureDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get Cpu architecture default response
func (o *GetCPUArchitectureDefault) Code() int {
	return o._statusCode
}

func (o *GetCPUArchitectureDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cpu-architectures/{cpu_architecture_id}][%d] getCpuArchitecture default  %+v", o._statusCode, o.Payload)
}
func (o *GetCPUArchitectureDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetCPUArchitectureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
