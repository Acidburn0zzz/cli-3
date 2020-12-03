// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IngredientVersionRevisionCore Ingredient Version Revision Core
//
// The fields of an ingredient version that can be updated by creating a new revision
//
// swagger:model ingredientVersionRevisionCore
type IngredientVersionRevisionCore struct {
	IngredientVersionRevisionCoreAllOf0

	Revision
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IngredientVersionRevisionCore) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IngredientVersionRevisionCoreAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IngredientVersionRevisionCoreAllOf0 = aO0

	// AO1
	var aO1 Revision
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Revision = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IngredientVersionRevisionCore) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.IngredientVersionRevisionCoreAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.Revision)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ingredient version revision core
func (m *IngredientVersionRevisionCore) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with IngredientVersionRevisionCoreAllOf0
	if err := m.IngredientVersionRevisionCoreAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with Revision
	if err := m.Revision.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionRevisionCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionRevisionCore) UnmarshalBinary(b []byte) error {
	var res IngredientVersionRevisionCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
