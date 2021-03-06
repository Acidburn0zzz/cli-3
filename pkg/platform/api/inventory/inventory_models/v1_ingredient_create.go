// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1IngredientCreate Ingredient Create
//
// A unique ingredient that can be used in a recipe. This ingredient model is only used during creation to allow creating ingredient versions along with the ingredient.
// swagger:model v1IngredientCreate
type V1IngredientCreate struct {
	V1IngredientCreateAllOf0

	V1IngredientCreateAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1IngredientCreate) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1IngredientCreateAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1IngredientCreateAllOf0 = aO0

	// AO1
	var aO1 V1IngredientCreateAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1IngredientCreateAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1IngredientCreate) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1IngredientCreateAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1IngredientCreateAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 ingredient create
func (m *V1IngredientCreate) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1IngredientCreateAllOf0
	if err := m.V1IngredientCreateAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1IngredientCreateAllOf1
	if err := m.V1IngredientCreateAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientCreate) UnmarshalBinary(b []byte) error {
	var res V1IngredientCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
