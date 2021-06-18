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

// GetIngredientOptionSetReader is a Reader for the GetIngredientOptionSet structure.
type GetIngredientOptionSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIngredientOptionSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIngredientOptionSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIngredientOptionSetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIngredientOptionSetOK creates a GetIngredientOptionSetOK with default headers values
func NewGetIngredientOptionSetOK() *GetIngredientOptionSetOK {
	return &GetIngredientOptionSetOK{}
}

/* GetIngredientOptionSetOK describes a response with status code 200, with default header values.

The retrieved ingredient option set
*/
type GetIngredientOptionSetOK struct {
	Payload *inventory_models.IngredientOptionSet
}

func (o *GetIngredientOptionSetOK) Error() string {
	return fmt.Sprintf("[GET /v1/ingredient-option-sets/{ingredient_option_set_id}][%d] getIngredientOptionSetOK  %+v", 200, o.Payload)
}
func (o *GetIngredientOptionSetOK) GetPayload() *inventory_models.IngredientOptionSet {
	return o.Payload
}

func (o *GetIngredientOptionSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.IngredientOptionSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIngredientOptionSetDefault creates a GetIngredientOptionSetDefault with default headers values
func NewGetIngredientOptionSetDefault(code int) *GetIngredientOptionSetDefault {
	return &GetIngredientOptionSetDefault{
		_statusCode: code,
	}
}

/* GetIngredientOptionSetDefault describes a response with status code -1, with default header values.

generic error response
*/
type GetIngredientOptionSetDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get ingredient option set default response
func (o *GetIngredientOptionSetDefault) Code() int {
	return o._statusCode
}

func (o *GetIngredientOptionSetDefault) Error() string {
	return fmt.Sprintf("[GET /v1/ingredient-option-sets/{ingredient_option_set_id}][%d] getIngredientOptionSet default  %+v", o._statusCode, o.Payload)
}
func (o *GetIngredientOptionSetDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetIngredientOptionSetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
