// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IngredientVersionAndUsageType Ingredient, Version, and Usage Type
//
// An ingredient, its version, and the way the version uses a particular ingredient option set
//
// swagger:model ingredientVersionAndUsageType
type IngredientVersionAndUsageType struct {
	IngredientAndVersion

	IngredientVersionAndUsageTypeAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IngredientVersionAndUsageType) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IngredientAndVersion
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IngredientAndVersion = aO0

	// AO1
	var aO1 IngredientVersionAndUsageTypeAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.IngredientVersionAndUsageTypeAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IngredientVersionAndUsageType) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.IngredientAndVersion)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.IngredientVersionAndUsageTypeAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ingredient version and usage type
func (m *IngredientVersionAndUsageType) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IngredientAndVersion
	if err := m.IngredientAndVersion.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with IngredientVersionAndUsageTypeAllOf1
	if err := m.IngredientVersionAndUsageTypeAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionAndUsageType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionAndUsageType) UnmarshalBinary(b []byte) error {
	var res IngredientVersionAndUsageType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
