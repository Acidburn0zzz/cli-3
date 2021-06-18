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

// NormalizeNamesReader is a Reader for the NormalizeNames structure.
type NormalizeNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NormalizeNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNormalizeNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNormalizeNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNormalizeNamesOK creates a NormalizeNamesOK with default headers values
func NewNormalizeNamesOK() *NormalizeNamesOK {
	return &NormalizeNamesOK{}
}

/* NormalizeNamesOK describes a response with status code 200, with default header values.

A list of mappings from requested name to normalized name.
*/
type NormalizeNamesOK struct {
	Payload *inventory_models.NormalizedNames
}

func (o *NormalizeNamesOK) Error() string {
	return fmt.Sprintf("[POST /v1/namespaces/normalized-names][%d] normalizeNamesOK  %+v", 200, o.Payload)
}
func (o *NormalizeNamesOK) GetPayload() *inventory_models.NormalizedNames {
	return o.Payload
}

func (o *NormalizeNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.NormalizedNames)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNormalizeNamesDefault creates a NormalizeNamesDefault with default headers values
func NewNormalizeNamesDefault(code int) *NormalizeNamesDefault {
	return &NormalizeNamesDefault{
		_statusCode: code,
	}
}

/* NormalizeNamesDefault describes a response with status code -1, with default header values.

generic error response
*/
type NormalizeNamesDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the normalize names default response
func (o *NormalizeNamesDefault) Code() int {
	return o._statusCode
}

func (o *NormalizeNamesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/namespaces/normalized-names][%d] normalizeNames default  %+v", o._statusCode, o.Payload)
}
func (o *NormalizeNamesDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *NormalizeNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
