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

// GpuArchitectureCoreAllOf0 gpu architecture core all of0
//
// swagger:model gpuArchitectureCoreAllOf0
type GpuArchitectureCoreAllOf0 struct {

	// The name of the GPU architecture
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this gpu architecture core all of0
func (m *GpuArchitectureCoreAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpuArchitectureCoreAllOf0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gpu architecture core all of0 based on context it is used
func (m *GpuArchitectureCoreAllOf0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GpuArchitectureCoreAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GpuArchitectureCoreAllOf0) UnmarshalBinary(b []byte) error {
	var res GpuArchitectureCoreAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
