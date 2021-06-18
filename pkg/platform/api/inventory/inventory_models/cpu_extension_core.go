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

// CPUExtensionCore CPU Extension Core
//
// The properties of a CPU extension needed to create a new one
//
// swagger:model cpuExtensionCore
type CPUExtensionCore struct {
	CPUExtensionCoreAllOf0

	RevisionedFeatureProvider
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CPUExtensionCore) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CPUExtensionCoreAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CPUExtensionCoreAllOf0 = aO0

	// AO1
	var aO1 RevisionedFeatureProvider
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.RevisionedFeatureProvider = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CPUExtensionCore) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CPUExtensionCoreAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.RevisionedFeatureProvider)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this cpu extension core
func (m *CPUExtensionCore) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CPUExtensionCoreAllOf0
	if err := m.CPUExtensionCoreAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RevisionedFeatureProvider
	if err := m.RevisionedFeatureProvider.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this cpu extension core based on the context it is used
func (m *CPUExtensionCore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CPUExtensionCoreAllOf0
	if err := m.CPUExtensionCoreAllOf0.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RevisionedFeatureProvider
	if err := m.RevisionedFeatureProvider.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CPUExtensionCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUExtensionCore) UnmarshalBinary(b []byte) error {
	var res CPUExtensionCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
