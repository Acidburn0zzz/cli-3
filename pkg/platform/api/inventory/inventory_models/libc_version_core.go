// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibcVersionCore Libc Version Core
//
// The properties of a libc version needed to create a new one
//
// swagger:model libcVersionCore
type LibcVersionCore struct {
	VersionInfo

	RevisionedFeatureProvider
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *LibcVersionCore) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VersionInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VersionInfo = aO0

	// AO1
	var aO1 RevisionedFeatureProvider
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.RevisionedFeatureProvider = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m LibcVersionCore) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.VersionInfo)
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

// Validate validates this libc version core
func (m *LibcVersionCore) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VersionInfo
	if err := m.VersionInfo.Validate(formats); err != nil {
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

// MarshalBinary interface implementation
func (m *LibcVersionCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LibcVersionCore) UnmarshalBinary(b []byte) error {
	var res LibcVersionCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
