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

// LibcAllOf0 libc all of0
//
// swagger:model libcAllOf0
type LibcAllOf0 struct {

	// libc id
	// Required: true
	// Format: uuid
	LibcID *strfmt.UUID `json:"libc_id"`

	// links
	// Required: true
	Links *SelfLink `json:"links"`
}

// Validate validates this libc all of0
func (m *LibcAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLibcID(formats); err != nil {
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

func (m *LibcAllOf0) validateLibcID(formats strfmt.Registry) error {

	if err := validate.Required("libc_id", "body", m.LibcID); err != nil {
		return err
	}

	if err := validate.FormatOf("libc_id", "body", "uuid", m.LibcID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LibcAllOf0) validateLinks(formats strfmt.Registry) error {

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

// ContextValidate validate this libc all of0 based on the context it is used
func (m *LibcAllOf0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LibcAllOf0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LibcAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LibcAllOf0) UnmarshalBinary(b []byte) error {
	var res LibcAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
