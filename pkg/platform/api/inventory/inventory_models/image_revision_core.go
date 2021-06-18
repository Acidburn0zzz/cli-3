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

// ImageRevisionCore Image Revision Core
//
// The properties of an image revision needed to create a new one
//
// swagger:model imageRevisionCore
type ImageRevisionCore struct {
	VersionInfo

	RevisionedFeatureProvider

	ImageRevisionCoreAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ImageRevisionCore) UnmarshalJSON(raw []byte) error {
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

	// AO2
	var aO2 ImageRevisionCoreAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.ImageRevisionCoreAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ImageRevisionCore) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

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

	aO2, err := swag.WriteJSON(m.ImageRevisionCoreAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this image revision core
func (m *ImageRevisionCore) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VersionInfo
	if err := m.VersionInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RevisionedFeatureProvider
	if err := m.RevisionedFeatureProvider.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ImageRevisionCoreAllOf2
	if err := m.ImageRevisionCoreAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this image revision core based on the context it is used
func (m *ImageRevisionCore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VersionInfo
	if err := m.VersionInfo.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RevisionedFeatureProvider
	if err := m.RevisionedFeatureProvider.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ImageRevisionCoreAllOf2
	if err := m.ImageRevisionCoreAllOf2.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ImageRevisionCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageRevisionCore) UnmarshalBinary(b []byte) error {
	var res ImageRevisionCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
