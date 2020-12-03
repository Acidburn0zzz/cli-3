// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KernelAllOf0 kernel all of0
//
// swagger:model kernelAllOf0
type KernelAllOf0 struct {

	// kernel id
	// Required: true
	// Format: uuid
	KernelID *strfmt.UUID `json:"kernel_id"`

	// links
	// Required: true
	Links *SelfLink `json:"links"`
}

// Validate validates this kernel all of0
func (m *KernelAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKernelID(formats); err != nil {
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

func (m *KernelAllOf0) validateKernelID(formats strfmt.Registry) error {

	if err := validate.Required("kernel_id", "body", m.KernelID); err != nil {
		return err
	}

	if err := validate.FormatOf("kernel_id", "body", "uuid", m.KernelID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KernelAllOf0) validateLinks(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *KernelAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KernelAllOf0) UnmarshalBinary(b []byte) error {
	var res KernelAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
