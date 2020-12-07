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

// AddIngredientOptionSetRevisionReader is a Reader for the AddIngredientOptionSetRevision structure.
type AddIngredientOptionSetRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddIngredientOptionSetRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddIngredientOptionSetRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddIngredientOptionSetRevisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddIngredientOptionSetRevisionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddIngredientOptionSetRevisionOK creates a AddIngredientOptionSetRevisionOK with default headers values
func NewAddIngredientOptionSetRevisionOK() *AddIngredientOptionSetRevisionOK {
	return &AddIngredientOptionSetRevisionOK{}
}

/*AddIngredientOptionSetRevisionOK handles this case with default header values.

The updated state of the ingredient option set
*/
type AddIngredientOptionSetRevisionOK struct {
	Payload *inventory_models.IngredientOptionSet
}

func (o *AddIngredientOptionSetRevisionOK) Error() string {
	return fmt.Sprintf("[POST /v1/ingredient-option-sets/{ingredient_option_set_id}/revisions][%d] addIngredientOptionSetRevisionOK  %+v", 200, o.Payload)
}

func (o *AddIngredientOptionSetRevisionOK) GetPayload() *inventory_models.IngredientOptionSet {
	return o.Payload
}

func (o *AddIngredientOptionSetRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.IngredientOptionSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddIngredientOptionSetRevisionBadRequest creates a AddIngredientOptionSetRevisionBadRequest with default headers values
func NewAddIngredientOptionSetRevisionBadRequest() *AddIngredientOptionSetRevisionBadRequest {
	return &AddIngredientOptionSetRevisionBadRequest{}
}

/*AddIngredientOptionSetRevisionBadRequest handles this case with default header values.

If the ingredient option set revision is invalid
*/
type AddIngredientOptionSetRevisionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddIngredientOptionSetRevisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/ingredient-option-sets/{ingredient_option_set_id}/revisions][%d] addIngredientOptionSetRevisionBadRequest  %+v", 400, o.Payload)
}

func (o *AddIngredientOptionSetRevisionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddIngredientOptionSetRevisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddIngredientOptionSetRevisionDefault creates a AddIngredientOptionSetRevisionDefault with default headers values
func NewAddIngredientOptionSetRevisionDefault(code int) *AddIngredientOptionSetRevisionDefault {
	return &AddIngredientOptionSetRevisionDefault{
		_statusCode: code,
	}
}

/*AddIngredientOptionSetRevisionDefault handles this case with default header values.

If there is an error processing the ingredient option set revision
*/
type AddIngredientOptionSetRevisionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add ingredient option set revision default response
func (o *AddIngredientOptionSetRevisionDefault) Code() int {
	return o._statusCode
}

func (o *AddIngredientOptionSetRevisionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ingredient-option-sets/{ingredient_option_set_id}/revisions][%d] addIngredientOptionSetRevision default  %+v", o._statusCode, o.Payload)
}

func (o *AddIngredientOptionSetRevisionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddIngredientOptionSetRevisionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}