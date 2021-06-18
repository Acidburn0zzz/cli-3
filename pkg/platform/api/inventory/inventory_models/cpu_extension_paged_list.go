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

// CPUExtensionPagedList CPU Extension Paged List
//
// A paginated list of CPU extensions
//
// swagger:model cpuExtensionPagedList
type CPUExtensionPagedList struct {

	// A page of CPU extensions
	// Required: true
	CPUExtensions []*CPUExtension `json:"cpu_extensions"`

	// links
	// Required: true
	Links *PagingLinks `json:"links"`

	// paging
	// Required: true
	Paging *Paging `json:"paging"`
}

// Validate validates this cpu extension paged list
func (m *CPUExtensionPagedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUExtensions(formats); err != nil {
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

func (m *CPUExtensionPagedList) validateCPUExtensions(formats strfmt.Registry) error {

	if err := validate.Required("cpu_extensions", "body", m.CPUExtensions); err != nil {
		return err
	}

	for i := 0; i < len(m.CPUExtensions); i++ {
		if swag.IsZero(m.CPUExtensions[i]) { // not required
			continue
		}

		if m.CPUExtensions[i] != nil {
			if err := m.CPUExtensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpu_extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CPUExtensionPagedList) validateLinks(formats strfmt.Registry) error {

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

func (m *CPUExtensionPagedList) validatePaging(formats strfmt.Registry) error {

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

// ContextValidate validate this cpu extension paged list based on the context it is used
func (m *CPUExtensionPagedList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPUExtensions(ctx, formats); err != nil {
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

func (m *CPUExtensionPagedList) contextValidateCPUExtensions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CPUExtensions); i++ {

		if m.CPUExtensions[i] != nil {
			if err := m.CPUExtensions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpu_extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CPUExtensionPagedList) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CPUExtensionPagedList) contextValidatePaging(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CPUExtensionPagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CPUExtensionPagedList) UnmarshalBinary(b []byte) error {
	var res CPUExtensionPagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
