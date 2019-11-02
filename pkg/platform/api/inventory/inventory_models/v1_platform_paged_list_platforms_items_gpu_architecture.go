// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1PlatformPagedListPlatformsItemsGpuArchitecture GPU Architecture
//
// The full GPU architecture data model
// swagger:model v1PlatformPagedListPlatformsItemsGpuArchitecture
type V1PlatformPagedListPlatformsItemsGpuArchitecture struct {
	V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf0

	V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf1

	V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1PlatformPagedListPlatformsItemsGpuArchitecture) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf0 = aO0

	// AO1
	var aO1 V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf1 = aO1

	// AO2
	var aO2 V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1PlatformPagedListPlatformsItemsGpuArchitecture) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 platform paged list platforms items gpu architecture
func (m *V1PlatformPagedListPlatformsItemsGpuArchitecture) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf0
	if err := m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf1
	if err := m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf2
	if err := m.V1PlatformPagedListPlatformsItemsGpuArchitectureAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1PlatformPagedListPlatformsItemsGpuArchitecture) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PlatformPagedListPlatformsItemsGpuArchitecture) UnmarshalBinary(b []byte) error {
	var res V1PlatformPagedListPlatformsItemsGpuArchitecture
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}