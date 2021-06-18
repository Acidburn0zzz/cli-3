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

// BuildFlagAllOf0 build flag all of0
//
// swagger:model buildFlagAllOf0
type BuildFlagAllOf0 struct {

	// build flag id
	// Required: true
	// Format: uuid
	BuildFlagID *strfmt.UUID `json:"build_flag_id"`

	// links
	// Required: true
	Links *SelfLink `json:"links"`
}

// Validate validates this build flag all of0
func (m *BuildFlagAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildFlagID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildFlagAllOf0) validateBuildFlagID(formats strfmt.Registry) error {

	if err := validate.Required("build_flag_id", "body", m.BuildFlagID); err != nil {
		return err
	}

	if err := validate.FormatOf("build_flag_id", "body", "uuid", m.BuildFlagID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildFlagAllOf0) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this build flag all of0 based on the context it is used
func (m *BuildFlagAllOf0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildFlagAllOf0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildFlagAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildFlagAllOf0) UnmarshalBinary(b []byte) error {
	var res BuildFlagAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
