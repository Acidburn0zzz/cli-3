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

// GetIngredientsReader is a Reader for the GetIngredients structure.
type GetIngredientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIngredientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIngredientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIngredientsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIngredientsOK creates a GetIngredientsOK with default headers values
func NewGetIngredientsOK() *GetIngredientsOK {
	return &GetIngredientsOK{}
}

/* GetIngredientsOK describes a response with status code 200, with default header values.

A paginated list of ingredients
*/
type GetIngredientsOK struct {
	Payload *inventory_models.IngredientPagedList
}

func (o *GetIngredientsOK) Error() string {
	return fmt.Sprintf("[GET /v1/ingredients][%d] getIngredientsOK  %+v", 200, o.Payload)
}
func (o *GetIngredientsOK) GetPayload() *inventory_models.IngredientPagedList {
	return o.Payload
}

func (o *GetIngredientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.IngredientPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIngredientsDefault creates a GetIngredientsDefault with default headers values
func NewGetIngredientsDefault(code int) *GetIngredientsDefault {
	return &GetIngredientsDefault{
		_statusCode: code,
	}
}

/* GetIngredientsDefault describes a response with status code -1, with default header values.

generic error response
*/
type GetIngredientsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get ingredients default response
func (o *GetIngredientsDefault) Code() int {
	return o._statusCode
}

func (o *GetIngredientsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/ingredients][%d] getIngredients default  %+v", o._statusCode, o.Payload)
}
func (o *GetIngredientsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetIngredientsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
