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

// BuildScript Build Script
//
// A short piece of scripting code that can be used to build an ingredient. This model contains all build script properties and is returned from read requests
//
// swagger:model buildScript
type BuildScript struct {
	BuildScriptAllOf0

	BuildScriptCore
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BuildScript) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuildScriptAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuildScriptAllOf0 = aO0

	// AO1
	var aO1 BuildScriptCore
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.BuildScriptCore = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BuildScript) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BuildScriptAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.BuildScriptCore)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this build script
func (m *BuildScript) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuildScriptAllOf0
	if err := m.BuildScriptAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BuildScriptCore
	if err := m.BuildScriptCore.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this build script based on the context it is used
func (m *BuildScript) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuildScriptAllOf0
	if err := m.BuildScriptAllOf0.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with BuildScriptCore
	if err := m.BuildScriptCore.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BuildScript) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildScript) UnmarshalBinary(b []byte) error {
	var res BuildScript
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
