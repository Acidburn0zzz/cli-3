// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NormalizedNames Unnormalized Names
//
// A list of mappings from unnormalized to normalized names.
//
// swagger:model normalizedNames
type NormalizedNames struct {

	// The normalized name mappings
	// Required: true
	NormalizedNames []*NormalizedName `json:"normalized_names"`
}

// Validate validates this normalized names
func (m *NormalizedNames) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNormalizedNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NormalizedNames) validateNormalizedNames(formats strfmt.Registry) error {

	if err := validate.Required("normalized_names", "body", m.NormalizedNames); err != nil {
		return err
	}

	for i := 0; i < len(m.NormalizedNames); i++ {
		if swag.IsZero(m.NormalizedNames[i]) { // not required
			continue
		}

		if m.NormalizedNames[i] != nil {
			if err := m.NormalizedNames[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("normalized_names" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this normalized names based on the context it is used
func (m *NormalizedNames) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNormalizedNames(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NormalizedNames) contextValidateNormalizedNames(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NormalizedNames); i++ {

		if m.NormalizedNames[i] != nil {
			if err := m.NormalizedNames[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("normalized_names" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NormalizedNames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NormalizedNames) UnmarshalBinary(b []byte) error {
	var res NormalizedNames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
