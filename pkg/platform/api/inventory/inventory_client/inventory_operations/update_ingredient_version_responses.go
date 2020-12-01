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

// UpdateIngredientVersionReader is a Reader for the UpdateIngredientVersion structure.
type UpdateIngredientVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIngredientVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIngredientVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateIngredientVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateIngredientVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateIngredientVersionOK creates a UpdateIngredientVersionOK with default headers values
func NewUpdateIngredientVersionOK() *UpdateIngredientVersionOK {
	return &UpdateIngredientVersionOK{}
}

/*UpdateIngredientVersionOK handles this case with default header values.

The updated state of the ingredient version
*/
type UpdateIngredientVersionOK struct {
	Payload *inventory_models.IngredientVersion
}

func (o *UpdateIngredientVersionOK) Error() string {
	return fmt.Sprintf("[PUT /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}][%d] updateIngredientVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateIngredientVersionOK) GetPayload() *inventory_models.IngredientVersion {
	return o.Payload
}

func (o *UpdateIngredientVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.IngredientVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIngredientVersionBadRequest creates a UpdateIngredientVersionBadRequest with default headers values
func NewUpdateIngredientVersionBadRequest() *UpdateIngredientVersionBadRequest {
	return &UpdateIngredientVersionBadRequest{}
}

/*UpdateIngredientVersionBadRequest handles this case with default header values.

If the ingredient version is invalid
*/
type UpdateIngredientVersionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *UpdateIngredientVersionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}][%d] updateIngredientVersionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIngredientVersionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *UpdateIngredientVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIngredientVersionDefault creates a UpdateIngredientVersionDefault with default headers values
func NewUpdateIngredientVersionDefault(code int) *UpdateIngredientVersionDefault {
	return &UpdateIngredientVersionDefault{
		_statusCode: code,
	}
}

/*UpdateIngredientVersionDefault handles this case with default header values.

If there is an error processing the request
*/
type UpdateIngredientVersionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the update ingredient version default response
func (o *UpdateIngredientVersionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIngredientVersionDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/ingredients/{ingredient_id}/versions/{ingredient_version_id}][%d] updateIngredientVersion default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateIngredientVersionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *UpdateIngredientVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
