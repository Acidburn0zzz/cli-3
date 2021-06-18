// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestAPIValidationError REST API validation error body
//
// A error for when invalid input is supplied to an endpoint.
//
// swagger:model restApiValidationError
type RestAPIValidationError struct {
	RestAPIError

	RestAPIValidationErrorAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RestAPIValidationError) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 RestAPIError
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.RestAPIError = aO0

	// AO1
	var aO1 RestAPIValidationErrorAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.RestAPIValidationErrorAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RestAPIValidationError) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.RestAPIError)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.RestAPIValidationErrorAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this rest Api validation error
func (m *RestAPIValidationError) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RestAPIError
	if err := m.RestAPIError.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RestAPIValidationErrorAllOf1
	if err := m.RestAPIValidationErrorAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this rest Api validation error based on the context it is used
func (m *RestAPIValidationError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with RestAPIError
	if err := m.RestAPIError.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RestAPIValidationErrorAllOf1
	if err := m.RestAPIValidationErrorAllOf1.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RestAPIValidationError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestAPIValidationError) UnmarshalBinary(b []byte) error {
	var res RestAPIValidationError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
