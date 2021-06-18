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

// GetNamespaceIngredientsReader is a Reader for the GetNamespaceIngredients structure.
type GetNamespaceIngredientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespaceIngredientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNamespaceIngredientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNamespaceIngredientsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNamespaceIngredientsOK creates a GetNamespaceIngredientsOK with default headers values
func NewGetNamespaceIngredientsOK() *GetNamespaceIngredientsOK {
	return &GetNamespaceIngredientsOK{}
}

/* GetNamespaceIngredientsOK describes a response with status code 200, with default header values.

A paginated list of ingredients and versions
*/
type GetNamespaceIngredientsOK struct {
	Payload *inventory_models.IngredientAndVersionPagedList
}

func (o *GetNamespaceIngredientsOK) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces/ingredients][%d] getNamespaceIngredientsOK  %+v", 200, o.Payload)
}
func (o *GetNamespaceIngredientsOK) GetPayload() *inventory_models.IngredientAndVersionPagedList {
	return o.Payload
}

func (o *GetNamespaceIngredientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.IngredientAndVersionPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceIngredientsDefault creates a GetNamespaceIngredientsDefault with default headers values
func NewGetNamespaceIngredientsDefault(code int) *GetNamespaceIngredientsDefault {
	return &GetNamespaceIngredientsDefault{
		_statusCode: code,
	}
}

/* GetNamespaceIngredientsDefault describes a response with status code -1, with default header values.

generic error response
*/
type GetNamespaceIngredientsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get namespace ingredients default response
func (o *GetNamespaceIngredientsDefault) Code() int {
	return o._statusCode
}

func (o *GetNamespaceIngredientsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces/ingredients][%d] getNamespaceIngredients default  %+v", o._statusCode, o.Payload)
}
func (o *GetNamespaceIngredientsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetNamespaceIngredientsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
