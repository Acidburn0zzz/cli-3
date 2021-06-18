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

// GpuArchitecturePagedList GPU Architecture Paged List
//
// A paginated list of GPU architectures
//
// swagger:model gpuArchitecturePagedList
type GpuArchitecturePagedList struct {

	// A page of GPU architectures
	// Required: true
	GpuArchitectures []*GpuArchitecture `json:"gpu_architectures"`

	// links
	// Required: true
	Links *PagingLinks `json:"links"`

	// paging
	// Required: true
	Paging *Paging `json:"paging"`
}

// Validate validates this gpu architecture paged list
func (m *GpuArchitecturePagedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGpuArchitectures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpuArchitecturePagedList) validateGpuArchitectures(formats strfmt.Registry) error {

	if err := validate.Required("gpu_architectures", "body", m.GpuArchitectures); err != nil {
		return err
	}

	for i := 0; i < len(m.GpuArchitectures); i++ {
		if swag.IsZero(m.GpuArchitectures[i]) { // not required
			continue
		}

		if m.GpuArchitectures[i] != nil {
			if err := m.GpuArchitectures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpu_architectures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GpuArchitecturePagedList) validateLinks(formats strfmt.Registry) error {

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

func (m *GpuArchitecturePagedList) validatePaging(formats strfmt.Registry) error {

	if err := validate.Required("paging", "body", m.Paging); err != nil {
		return err
	}

	if m.Paging != nil {
		if err := m.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gpu architecture paged list based on the context it is used
func (m *GpuArchitecturePagedList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGpuArchitectures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaging(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GpuArchitecturePagedList) contextValidateGpuArchitectures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GpuArchitectures); i++ {

		if m.GpuArchitectures[i] != nil {
			if err := m.GpuArchitectures[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpu_architectures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GpuArchitecturePagedList) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GpuArchitecturePagedList) contextValidatePaging(ctx context.Context, formats strfmt.Registry) error {

	if m.Paging != nil {
		if err := m.Paging.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GpuArchitecturePagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GpuArchitecturePagedList) UnmarshalBinary(b []byte) error {
	var res GpuArchitecturePagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
