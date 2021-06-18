// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestAPIValidationErrorItem REST API validation error body item
//
// swagger:model restApiValidationErrorItem
type RestAPIValidationErrorItem struct {

	// Human-readable message about this specific validation error
	// Required: true
	Error *string `json:"error"`

	// A JSON path to the specific field that was invalid, if applicable
	JSONPath string `json:"jsonPath,omitempty"`
}

// Validate validates this rest Api validation error item
func (m *RestAPIValidationErrorItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestAPIValidationErrorItem) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rest Api validation error item based on context it is used
func (m *RestAPIValidationErrorItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestAPIValidationErrorItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestAPIValidationErrorItem) UnmarshalBinary(b []byte) error {
	var res RestAPIValidationErrorItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
