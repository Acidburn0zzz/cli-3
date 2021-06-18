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

// GetCPUArchitecturesReader is a Reader for the GetCPUArchitectures structure.
type GetCPUArchitecturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCPUArchitecturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCPUArchitecturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCPUArchitecturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCPUArchitecturesOK creates a GetCPUArchitecturesOK with default headers values
func NewGetCPUArchitecturesOK() *GetCPUArchitecturesOK {
	return &GetCPUArchitecturesOK{}
}

/* GetCPUArchitecturesOK describes a response with status code 200, with default header values.

A paginated list of CPU architectures
*/
type GetCPUArchitecturesOK struct {
	Payload *inventory_models.CPUArchitecturePagedList
}

func (o *GetCPUArchitecturesOK) Error() string {
	return fmt.Sprintf("[GET /v1/cpu-architectures][%d] getCpuArchitecturesOK  %+v", 200, o.Payload)
}
func (o *GetCPUArchitecturesOK) GetPayload() *inventory_models.CPUArchitecturePagedList {
	return o.Payload
}

func (o *GetCPUArchitecturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.CPUArchitecturePagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCPUArchitecturesDefault creates a GetCPUArchitecturesDefault with default headers values
func NewGetCPUArchitecturesDefault(code int) *GetCPUArchitecturesDefault {
	return &GetCPUArchitecturesDefault{
		_statusCode: code,
	}
}

/* GetCPUArchitecturesDefault describes a response with status code -1, with default header values.

generic error response
*/
type GetCPUArchitecturesDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get Cpu architectures default response
func (o *GetCPUArchitecturesDefault) Code() int {
	return o._statusCode
}

func (o *GetCPUArchitecturesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/cpu-architectures][%d] getCpuArchitectures default  %+v", o._statusCode, o.Payload)
}
func (o *GetCPUArchitecturesDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetCPUArchitecturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
